// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// middleware is an autogenerated mock type for the middleware type
type middleware struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *middleware) Execute(_a0 http.Handler) http.Handler {
	ret := _m.Called(_a0)

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func(http.Handler) http.Handler); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}
