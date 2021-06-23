// Copyright 2020 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package inv

import (
	"context"
	"crypto/md5"
	"fmt"

	"github.com/pkg/errors"

	"github.com/mendersoftware/inventory/model"
	"github.com/mendersoftware/inventory/store"
	"github.com/mendersoftware/inventory/store/mongo"
)

var (
	ErrETagsDontMatch       = errors.New("Received and saved ETags don't match")
	ErrMissingIfMatchHeader = errors.New("If-Match header is required")
)

// this inventory service interface
//go:generate ../utils/mockgen.sh
type InventoryApp interface {
	HealthCheck(ctx context.Context) error
	ListDevices(ctx context.Context, q store.ListQuery) ([]model.Device, int, error)
	GetDevice(ctx context.Context, id model.DeviceID) (*model.Device, error)
	AddDevice(ctx context.Context, d *model.Device) error
	UpsertAttributes(ctx context.Context, id model.DeviceID, attrs model.DeviceAttributes) error
	UpsertAttributesWithUpdated(ctx context.Context, id model.DeviceID, attrs model.DeviceAttributes, scope string) error
	UpsertDevicesStatuses(ctx context.Context, devices []model.DeviceUpdate, attrs model.DeviceAttributes) (*model.UpdateResult, error)
	ReplaceAttributes(ctx context.Context, id model.DeviceID, upsertAttrs model.DeviceAttributes, scope string) error
	GetFiltersAttributes(ctx context.Context) ([]model.FilterAttribute, error)
	UnsetDeviceGroup(ctx context.Context, id model.DeviceID, groupName model.GroupName) error
	UnsetDevicesGroup(
		ctx context.Context,
		deviceIDs []model.DeviceID,
		groupName model.GroupName,
	) (*model.UpdateResult, error)
	UpdateDeviceGroup(ctx context.Context, id model.DeviceID, group model.GroupName) error
	UpdateDevicesGroup(
		ctx context.Context,
		ids []model.DeviceID,
		group model.GroupName,
	) (*model.UpdateResult, error)
	ListGroups(ctx context.Context, filters []model.FilterPredicate) ([]model.GroupName, error)
	ListDevicesByGroup(ctx context.Context, group model.GroupName, skip int, limit int) ([]model.DeviceID, int, error)
	GetDeviceGroup(ctx context.Context, id model.DeviceID) (model.GroupName, error)
	DeleteDevice(ctx context.Context, id model.DeviceID) error
	DeleteDevices(
		ctx context.Context,
		ids []model.DeviceID,
	) (*model.UpdateResult, error)
	CreateTenant(ctx context.Context, tenant model.NewTenant) error
	SearchDevices(ctx context.Context, searchParams model.SearchParams) ([]model.Device, int, error)
}

type inventory struct {
	db store.DataStore
}

func NewInventory(d store.DataStore) InventoryApp {
	return &inventory{db: d}
}

func (i *inventory) HealthCheck(ctx context.Context) error {
	err := i.db.Ping(ctx)
	if err != nil {
		return errors.Wrap(err, "error reaching MongoDB")
	}
	return nil
}

func (i *inventory) ListDevices(ctx context.Context, q store.ListQuery) ([]model.Device, int, error) {
	devs, totalCount, err := i.db.GetDevices(ctx, q)

	if err != nil {
		return nil, -1, errors.Wrap(err, "failed to fetch devices")
	}

	return devs, totalCount, nil
}

func (i *inventory) GetDevice(ctx context.Context, id model.DeviceID) (*model.Device, error) {
	dev, err := i.db.GetDevice(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch device")
	}
	return dev, nil
}

func (i *inventory) AddDevice(ctx context.Context, dev *model.Device) error {
	if dev == nil {
		return errors.New("no device given")
	}
	err := i.db.AddDevice(ctx, dev)
	if err != nil {
		return errors.Wrap(err, "failed to add device")
	}
	return nil
}

func (i *inventory) DeleteDevices(
	ctx context.Context,
	ids []model.DeviceID,
) (*model.UpdateResult, error) {
	return i.db.DeleteDevices(ctx, ids)
}

func (i *inventory) DeleteDevice(ctx context.Context, id model.DeviceID) error {
	res, err := i.db.DeleteDevices(ctx, []model.DeviceID{id})
	if err != nil {
		return errors.Wrap(err, "failed to delete device")
	} else if res.DeletedCount < 1 {
		return store.ErrDevNotFound
	}
	return nil
}

func (i *inventory) UpsertAttributes(ctx context.Context, id model.DeviceID, attrs model.DeviceAttributes) error {
	if _, err := i.db.UpsertDevicesAttributes(
		ctx, []model.DeviceID{id}, attrs,
	); err != nil {
		return errors.Wrap(err, "failed to upsert attributes in db")
	}
	return nil
}

func (i *inventory) UpsertAttributesWithUpdated(ctx context.Context, id model.DeviceID, attrs model.DeviceAttributes, scope string) error {
	if scope == model.AttrScopeTags {
		device, err := i.db.GetDevice(ctx, id)
		if err != nil && err != store.ErrDevNotFound {
			return errors.Wrap(err, "failed to get the device")
		}
		// check provided ETag
		err = checkTagAttributesETag(ctx, device.Tags_etag)
		if err != nil {
			return err
		}
		// make all attributes list
		totalAttrs := model.DeviceAttributes{}
		attrsUpdated := make([]bool, len(attrs))
		for _, da := range device.Attributes {
			if da.Scope == model.AttrScopeTags {
				totalAttrs = append(totalAttrs, da)
				for i, a := range attrs {
					if da.Name == a.Name {
						totalAttrs[len(totalAttrs)-1] = a
						attrsUpdated[i] = true
						break
					}
				}
			}
		}
		for i, val := range attrsUpdated {
			if !val {
				totalAttrs = append(totalAttrs, attrs[i])
			}
		}
		// calculate new ETag
		ctx = calcTagAttributesETag(ctx, totalAttrs)
	}

	if _, err := i.db.UpsertDevicesAttributesWithUpdated(
		ctx, []model.DeviceID{id}, attrs,
	); err != nil {
		return errors.Wrap(err, "failed to upsert attributes in db")
	}
	return nil
}

func (i *inventory) ReplaceAttributes(ctx context.Context, id model.DeviceID, upsertAttrs model.DeviceAttributes, scope string) error {
	device, err := i.db.GetDevice(ctx, id)
	if err != nil && err != store.ErrDevNotFound {
		return errors.Wrap(err, "failed to get the device")
	}
	removeAttrs := model.DeviceAttributes{}
	if device != nil {
		for _, attr := range device.Attributes {
			if attr.Scope == scope {
				update := false
				for _, upsertAttr := range upsertAttrs {
					if upsertAttr.Name == attr.Name {
						update = true
						break
					}
				}
				if !update {
					removeAttrs = append(removeAttrs, attr)
				}
			}
		}
	}

	if scope == model.AttrScopeTags {
		err := checkTagAttributesETag(ctx, device.Tags_etag)
		if err != nil {
			return err
		}
		ctx = calcTagAttributesETag(ctx, upsertAttrs)
	}

	if _, err := i.db.UpsertRemoveDeviceAttributes(ctx, id, upsertAttrs, removeAttrs); err != nil {
		return errors.Wrap(err, "failed to replace attributes in db")
	}
	return nil
}

func checkTagAttributesETag(ctx context.Context, etag *string) error {
	ifMatchHeader := ctx.Value(model.CtxKeyIfMatchHeader)
	if etag != nil && *etag != "" {
		if ifMatchHeader == "" {
			return ErrMissingIfMatchHeader
		}
		if *etag != ifMatchHeader {
			return ErrETagsDontMatch
		}
	} else {
		if ifMatchHeader != "" {
			return ErrETagsDontMatch
		}
	}
	return nil
}

// calculates ETag for attributes with 'tags' scope
// ETag is calculated as MD5 checksum of a sum of all timestamps in Unix format
func calcTagAttributesETag(ctx context.Context, attrs model.DeviceAttributes) context.Context {
	etag := ""
	if len(attrs) > 0 {
		var sum int64
		for _, attr := range attrs {
			if attr.Scope == model.AttrScopeTags {
				sum += attr.Timestamp.Unix()
			}
		}
		if sum > 0 {
			data := []byte(fmt.Sprint(sum))
			etag = fmt.Sprintf("%x", md5.Sum(data))
		}
	}
	ctx = context.WithValue(ctx, model.CtxKeyETag, etag)
	return ctx
}

func (i *inventory) GetFiltersAttributes(ctx context.Context) ([]model.FilterAttribute, error) {
	attributes, err := i.db.GetFiltersAttributes(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get filter attributes from the db")
	}
	return attributes, nil
}

func (i *inventory) UpsertDevicesStatuses(
	ctx context.Context,
	devices []model.DeviceUpdate,
	attrs model.DeviceAttributes,
) (*model.UpdateResult, error) {
	return i.db.UpsertDevicesAttributesWithRevision(ctx, devices, attrs)
}

func (i *inventory) UnsetDevicesGroup(
	ctx context.Context,
	deviceIDs []model.DeviceID,
	groupName model.GroupName,
) (*model.UpdateResult, error) {
	return i.db.UnsetDevicesGroup(ctx, deviceIDs, groupName)
}

func (i *inventory) UnsetDeviceGroup(ctx context.Context, id model.DeviceID, group model.GroupName) error {
	result, err := i.db.UnsetDevicesGroup(ctx, []model.DeviceID{id}, group)
	if err != nil {
		return errors.Wrap(err, "failed to unassign group from device")
	} else if result.MatchedCount <= 0 {
		return store.ErrDevNotFound
	}
	return nil
}

func (i *inventory) UpdateDevicesGroup(
	ctx context.Context,
	deviceIDs []model.DeviceID,
	group model.GroupName,
) (*model.UpdateResult, error) {
	return i.db.UpdateDevicesGroup(ctx, deviceIDs, group)
}

func (i *inventory) UpdateDeviceGroup(
	ctx context.Context,
	devid model.DeviceID,
	group model.GroupName,
) error {
	result, err := i.db.UpdateDevicesGroup(
		ctx, []model.DeviceID{devid}, group,
	)
	if err != nil {
		return errors.Wrap(err, "failed to add device to group")
	} else if result.MatchedCount <= 0 {
		return store.ErrDevNotFound
	}
	return nil
}

func (i *inventory) ListGroups(
	ctx context.Context,
	filters []model.FilterPredicate,
) ([]model.GroupName, error) {
	groups, err := i.db.ListGroups(ctx, filters)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list groups")
	}

	if groups == nil {
		return []model.GroupName{}, nil
	}
	return groups, nil
}

func (i *inventory) ListDevicesByGroup(ctx context.Context, group model.GroupName, skip, limit int) ([]model.DeviceID, int, error) {
	ids, totalCount, err := i.db.GetDevicesByGroup(ctx, group, skip, limit)
	if err != nil {
		if err == store.ErrGroupNotFound {
			return nil, -1, err
		} else {
			return nil, -1, errors.Wrap(err, "failed to list devices by group")
		}
	}

	return ids, totalCount, nil
}

func (i *inventory) GetDeviceGroup(ctx context.Context, id model.DeviceID) (model.GroupName, error) {
	group, err := i.db.GetDeviceGroup(ctx, id)
	if err != nil {
		if err == store.ErrDevNotFound {
			return "", err
		} else {
			return "", errors.Wrap(err, "failed to get device's group")
		}
	}

	return group, nil
}

func (i *inventory) CreateTenant(ctx context.Context, tenant model.NewTenant) error {
	if err := i.db.WithAutomigrate().
		MigrateTenant(ctx, mongo.DbVersion, tenant.ID); err != nil {
		return errors.Wrapf(err, "failed to apply migrations for tenant %v", tenant.ID)
	}
	return nil
}

func (i *inventory) SearchDevices(ctx context.Context, searchParams model.SearchParams) ([]model.Device, int, error) {
	devs, totalCount, err := i.db.SearchDevices(ctx, searchParams)

	if err != nil {
		return nil, -1, errors.Wrap(err, "failed to fetch devices")
	}

	return devs, totalCount, nil
}
