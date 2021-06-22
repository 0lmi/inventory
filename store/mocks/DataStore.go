// Code generated by mockery (devel). DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/mendersoftware/inventory/model"
	mock "github.com/stretchr/testify/mock"

	store "github.com/mendersoftware/inventory/store"
)

// DataStore is an autogenerated mock type for the DataStore type
type DataStore struct {
	mock.Mock
}

// AddDevice provides a mock function with given fields: ctx, dev
func (_m *DataStore) AddDevice(ctx context.Context, dev *model.Device) error {
	ret := _m.Called(ctx, dev)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Device) error); ok {
		r0 = rf(ctx, dev)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDevices provides a mock function with given fields: ctx, ids
func (_m *DataStore) DeleteDevices(ctx context.Context, ids []model.DeviceID) (*model.UpdateResult, error) {
	ret := _m.Called(ctx, ids)

	var r0 *model.UpdateResult
	if rf, ok := ret.Get(0).(func(context.Context, []model.DeviceID) *model.UpdateResult); ok {
		r0 = rf(ctx, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UpdateResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []model.DeviceID) error); ok {
		r1 = rf(ctx, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllAttributeNames provides a mock function with given fields: ctx
func (_m *DataStore) GetAllAttributeNames(ctx context.Context) ([]string, error) {
	ret := _m.Called(ctx)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDevice provides a mock function with given fields: ctx, id
func (_m *DataStore) GetDevice(ctx context.Context, id model.DeviceID) (*model.Device, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.Device
	if rf, ok := ret.Get(0).(func(context.Context, model.DeviceID) *model.Device); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.DeviceID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceGroup provides a mock function with given fields: ctx, id
func (_m *DataStore) GetDeviceGroup(ctx context.Context, id model.DeviceID) (model.GroupName, error) {
	ret := _m.Called(ctx, id)

	var r0 model.GroupName
	if rf, ok := ret.Get(0).(func(context.Context, model.DeviceID) model.GroupName); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.GroupName)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.DeviceID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDevices provides a mock function with given fields: ctx, q
func (_m *DataStore) GetDevices(ctx context.Context, q store.ListQuery) ([]model.Device, int, error) {
	ret := _m.Called(ctx, q)

	var r0 []model.Device
	if rf, ok := ret.Get(0).(func(context.Context, store.ListQuery) []model.Device); ok {
		r0 = rf(ctx, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Device)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, store.ListQuery) int); ok {
		r1 = rf(ctx, q)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, store.ListQuery) error); ok {
		r2 = rf(ctx, q)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetDevicesByGroup provides a mock function with given fields: ctx, group, skip, limit
func (_m *DataStore) GetDevicesByGroup(ctx context.Context, group model.GroupName, skip int, limit int) ([]model.DeviceID, int, error) {
	ret := _m.Called(ctx, group, skip, limit)

	var r0 []model.DeviceID
	if rf, ok := ret.Get(0).(func(context.Context, model.GroupName, int, int) []model.DeviceID); ok {
		r0 = rf(ctx, group, skip, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.DeviceID)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, model.GroupName, int, int) int); ok {
		r1 = rf(ctx, group, skip, limit)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, model.GroupName, int, int) error); ok {
		r2 = rf(ctx, group, skip, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetFiltersAttributes provides a mock function with given fields: ctx
func (_m *DataStore) GetFiltersAttributes(ctx context.Context) ([]model.FilterAttribute, error) {
	ret := _m.Called(ctx)

	var r0 []model.FilterAttribute
	if rf, ok := ret.Get(0).(func(context.Context) []model.FilterAttribute); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.FilterAttribute)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListGroups provides a mock function with given fields: ctx, filters
func (_m *DataStore) ListGroups(ctx context.Context, filters []model.FilterPredicate) ([]model.GroupName, error) {
	ret := _m.Called(ctx, filters)

	var r0 []model.GroupName
	if rf, ok := ret.Get(0).(func(context.Context, []model.FilterPredicate) []model.GroupName); ok {
		r0 = rf(ctx, filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.GroupName)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []model.FilterPredicate) error); ok {
		r1 = rf(ctx, filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Maintenance provides a mock function with given fields: ctx, version, tenantIDs
func (_m *DataStore) Maintenance(ctx context.Context, version string, tenantIDs ...string) error {
	_va := make([]interface{}, len(tenantIDs))
	for _i := range tenantIDs {
		_va[_i] = tenantIDs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, version)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...string) error); ok {
		r0 = rf(ctx, version, tenantIDs...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Migrate provides a mock function with given fields: ctx, version
func (_m *DataStore) Migrate(ctx context.Context, version string) error {
	ret := _m.Called(ctx, version)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, version)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MigrateTenant provides a mock function with given fields: ctx, version, tenantId
func (_m *DataStore) MigrateTenant(ctx context.Context, version string, tenantId string) error {
	ret := _m.Called(ctx, version, tenantId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, version, tenantId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Ping provides a mock function with given fields: ctx
func (_m *DataStore) Ping(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SearchDevices provides a mock function with given fields: ctx, searchParams
func (_m *DataStore) SearchDevices(ctx context.Context, searchParams model.SearchParams) ([]model.Device, int, error) {
	ret := _m.Called(ctx, searchParams)

	var r0 []model.Device
	if rf, ok := ret.Get(0).(func(context.Context, model.SearchParams) []model.Device); ok {
		r0 = rf(ctx, searchParams)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Device)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, model.SearchParams) int); ok {
		r1 = rf(ctx, searchParams)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, model.SearchParams) error); ok {
		r2 = rf(ctx, searchParams)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UnsetDevicesGroup provides a mock function with given fields: ctx, deviceIDs, group
func (_m *DataStore) UnsetDevicesGroup(ctx context.Context, deviceIDs []model.DeviceID, group model.GroupName) (*model.UpdateResult, error) {
	ret := _m.Called(ctx, deviceIDs, group)

	var r0 *model.UpdateResult
	if rf, ok := ret.Get(0).(func(context.Context, []model.DeviceID, model.GroupName) *model.UpdateResult); ok {
		r0 = rf(ctx, deviceIDs, group)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UpdateResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []model.DeviceID, model.GroupName) error); ok {
		r1 = rf(ctx, deviceIDs, group)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDevicesGroup provides a mock function with given fields: ctx, devIDs, group
func (_m *DataStore) UpdateDevicesGroup(ctx context.Context, devIDs []model.DeviceID, group model.GroupName) (*model.UpdateResult, error) {
	ret := _m.Called(ctx, devIDs, group)

	var r0 *model.UpdateResult
	if rf, ok := ret.Get(0).(func(context.Context, []model.DeviceID, model.GroupName) *model.UpdateResult); ok {
		r0 = rf(ctx, devIDs, group)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UpdateResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []model.DeviceID, model.GroupName) error); ok {
		r1 = rf(ctx, devIDs, group)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpsertDevicesAttributes provides a mock function with given fields: ctx, ids, attrs
func (_m *DataStore) UpsertDevicesAttributes(ctx context.Context, ids []model.DeviceID, attrs model.DeviceAttributes) (*model.UpdateResult, error) {
	ret := _m.Called(ctx, ids, attrs)

	var r0 *model.UpdateResult
	if rf, ok := ret.Get(0).(func(context.Context, []model.DeviceID, model.DeviceAttributes) *model.UpdateResult); ok {
		r0 = rf(ctx, ids, attrs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UpdateResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []model.DeviceID, model.DeviceAttributes) error); ok {
		r1 = rf(ctx, ids, attrs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpsertDevicesAttributesWithRevision provides a mock function with given fields: ctx, ids, attrs
func (_m *DataStore) UpsertDevicesAttributesWithRevision(ctx context.Context, ids []model.DeviceUpdate, attrs model.DeviceAttributes) (*model.UpdateResult, error) {
	ret := _m.Called(ctx, ids, attrs)

	var r0 *model.UpdateResult
	if rf, ok := ret.Get(0).(func(context.Context, []model.DeviceUpdate, model.DeviceAttributes) *model.UpdateResult); ok {
		r0 = rf(ctx, ids, attrs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UpdateResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []model.DeviceUpdate, model.DeviceAttributes) error); ok {
		r1 = rf(ctx, ids, attrs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpsertDevicesAttributesWithUpdated provides a mock function with given fields: ctx, ids, attrs
func (_m *DataStore) UpsertDevicesAttributesWithUpdated(ctx context.Context, ids []model.DeviceID, attrs model.DeviceAttributes) (*model.UpdateResult, error) {
	ret := _m.Called(ctx, ids, attrs)

	var r0 *model.UpdateResult
	if rf, ok := ret.Get(0).(func(context.Context, []model.DeviceID, model.DeviceAttributes) *model.UpdateResult); ok {
		r0 = rf(ctx, ids, attrs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UpdateResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []model.DeviceID, model.DeviceAttributes) error); ok {
		r1 = rf(ctx, ids, attrs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpsertRemoveDeviceAttributes provides a mock function with given fields: ctx, id, updateAttrs, removeAttrs
func (_m *DataStore) UpsertRemoveDeviceAttributes(ctx context.Context, id model.DeviceID, updateAttrs model.DeviceAttributes, removeAttrs model.DeviceAttributes) (*model.UpdateResult, error) {
	ret := _m.Called(ctx, id, updateAttrs, removeAttrs)

	var r0 *model.UpdateResult
	if rf, ok := ret.Get(0).(func(context.Context, model.DeviceID, model.DeviceAttributes, model.DeviceAttributes) *model.UpdateResult); ok {
		r0 = rf(ctx, id, updateAttrs, removeAttrs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UpdateResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.DeviceID, model.DeviceAttributes, model.DeviceAttributes) error); ok {
		r1 = rf(ctx, id, updateAttrs, removeAttrs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WithAutomigrate provides a mock function with given fields:
func (_m *DataStore) WithAutomigrate() store.DataStore {
	ret := _m.Called()

	var r0 store.DataStore
	if rf, ok := ret.Get(0).(func() store.DataStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.DataStore)
		}
	}

	return r0
}
