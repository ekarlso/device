// Code generated by mockery v2.10.0. DO NOT EDIT.

package wireguard

import mock "github.com/stretchr/testify/mock"

// MockNetworkConfigurer is an autogenerated mock type for the NetworkConfigurer type
type MockNetworkConfigurer struct {
	mock.Mock
}

// ApplyWireGuardConfig provides a mock function with given fields: peers
func (_m *MockNetworkConfigurer) ApplyWireGuardConfig(peers []Peer) error {
	ret := _m.Called(peers)

	var r0 error
	if rf, ok := ret.Get(0).(func([]Peer) error); ok {
		r0 = rf(peers)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConnectedDeviceCount provides a mock function with given fields:
func (_m *MockNetworkConfigurer) ConnectedDeviceCount() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ForwardRoutes provides a mock function with given fields: routes
func (_m *MockNetworkConfigurer) ForwardRoutes(routes []string) error {
	ret := _m.Called(routes)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(routes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetupIPTables provides a mock function with given fields:
func (_m *MockNetworkConfigurer) SetupIPTables() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetupInterface provides a mock function with given fields:
func (_m *MockNetworkConfigurer) SetupInterface() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
