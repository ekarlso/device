// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UrlOpener is an autogenerated mock type for the UrlOpener type
type UrlOpener struct {
	mock.Mock
}

// Execute provides a mock function with given fields:
func (_m *UrlOpener) Execute() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
