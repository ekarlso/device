// Code generated by mockery v2.10.0. DO NOT EDIT.

package pb

import (
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// MockPeer is an autogenerated mock type for the Peer type
type MockPeer struct {
	mock.Mock
}

// WritePeerConfig provides a mock function with given fields: _a0
func (_m *MockPeer) WritePeerConfig(_a0 io.Writer) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(io.Writer) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
