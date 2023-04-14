// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	interfaces "github.com/edgexfoundry/device-sdk-go/v3/pkg/interfaces"
	mock "github.com/stretchr/testify/mock"

	models "github.com/edgexfoundry/go-mod-core-contracts/v3/models"

	pkgmodels "github.com/edgexfoundry/device-sdk-go/v3/pkg/models"
)

// ProtocolDriver is an autogenerated mock type for the ProtocolDriver type
type ProtocolDriver struct {
	mock.Mock
}

// AddDevice provides a mock function with given fields: deviceName, protocols, adminState
func (_m *ProtocolDriver) AddDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) error {
	ret := _m.Called(deviceName, protocols, adminState)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, map[string]models.ProtocolProperties, models.AdminState) error); ok {
		r0 = rf(deviceName, protocols, adminState)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Discover provides a mock function with given fields:
func (_m *ProtocolDriver) Discover() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HandleReadCommands provides a mock function with given fields: deviceName, protocols, reqs
func (_m *ProtocolDriver) HandleReadCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []pkgmodels.CommandRequest) ([]*pkgmodels.CommandValue, error) {
	ret := _m.Called(deviceName, protocols, reqs)

	var r0 []*pkgmodels.CommandValue
	var r1 error
	if rf, ok := ret.Get(0).(func(string, map[string]models.ProtocolProperties, []pkgmodels.CommandRequest) ([]*pkgmodels.CommandValue, error)); ok {
		return rf(deviceName, protocols, reqs)
	}
	if rf, ok := ret.Get(0).(func(string, map[string]models.ProtocolProperties, []pkgmodels.CommandRequest) []*pkgmodels.CommandValue); ok {
		r0 = rf(deviceName, protocols, reqs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pkgmodels.CommandValue)
		}
	}

	if rf, ok := ret.Get(1).(func(string, map[string]models.ProtocolProperties, []pkgmodels.CommandRequest) error); ok {
		r1 = rf(deviceName, protocols, reqs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HandleWriteCommands provides a mock function with given fields: deviceName, protocols, reqs, params
func (_m *ProtocolDriver) HandleWriteCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []pkgmodels.CommandRequest, params []*pkgmodels.CommandValue) error {
	ret := _m.Called(deviceName, protocols, reqs, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, map[string]models.ProtocolProperties, []pkgmodels.CommandRequest, []*pkgmodels.CommandValue) error); ok {
		r0 = rf(deviceName, protocols, reqs, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Initialize provides a mock function with given fields: sdk
func (_m *ProtocolDriver) Initialize(sdk interfaces.DeviceServiceSDK) error {
	ret := _m.Called(sdk)

	var r0 error
	if rf, ok := ret.Get(0).(func(interfaces.DeviceServiceSDK) error); ok {
		r0 = rf(sdk)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveDevice provides a mock function with given fields: deviceName, protocols
func (_m *ProtocolDriver) RemoveDevice(deviceName string, protocols map[string]models.ProtocolProperties) error {
	ret := _m.Called(deviceName, protocols)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, map[string]models.ProtocolProperties) error); ok {
		r0 = rf(deviceName, protocols)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields:
func (_m *ProtocolDriver) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields: force
func (_m *ProtocolDriver) Stop(force bool) error {
	ret := _m.Called(force)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(force)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDevice provides a mock function with given fields: deviceName, protocols, adminState
func (_m *ProtocolDriver) UpdateDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) error {
	ret := _m.Called(deviceName, protocols, adminState)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, map[string]models.ProtocolProperties, models.AdminState) error); ok {
		r0 = rf(deviceName, protocols, adminState)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateDevice provides a mock function with given fields: device
func (_m *ProtocolDriver) ValidateDevice(device models.Device) error {
	ret := _m.Called(device)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Device) error); ok {
		r0 = rf(device)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewProtocolDriver interface {
	mock.TestingT
	Cleanup(func())
}

// NewProtocolDriver creates a new instance of ProtocolDriver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProtocolDriver(t mockConstructorTestingTNewProtocolDriver) *ProtocolDriver {
	mock := &ProtocolDriver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
