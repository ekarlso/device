// Code generated by mockery v2.10.0. DO NOT EDIT.

package notify

import mock "github.com/stretchr/testify/mock"

// mockLogFunc is an autogenerated mock type for the logFunc type
type mockLogFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0, _a1
func (_m *mockLogFunc) Execute(_a0 string, _a1 ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _a1...)
	_m.Called(_ca...)
}
