// Code generated by mockery v2.9.4. DO NOT EDIT.

package auth

import (
	context "context"
	http "net/http"

	mock "github.com/stretchr/testify/mock"

	pb "github.com/nais/device/pkg/pb"
)

// MockAuthenticator is an autogenerated mock type for the Authenticator type
type MockAuthenticator struct {
	mock.Mock
}

// AuthURL provides a mock function with given fields: w, r
func (_m *MockAuthenticator) AuthURL(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// Login provides a mock function with given fields: ctx, token, serial, platform
func (_m *MockAuthenticator) Login(ctx context.Context, token string, serial string, platform string) (*pb.Session, error) {
	ret := _m.Called(ctx, token, serial, platform)

	var r0 *pb.Session
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *pb.Session); ok {
		r0 = rf(ctx, token, serial, platform)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Session)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, token, serial, platform)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginHTTP provides a mock function with given fields: w, r
func (_m *MockAuthenticator) LoginHTTP(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// Validator provides a mock function with given fields:
func (_m *MockAuthenticator) Validator() func(http.Handler) http.Handler {
	ret := _m.Called()

	var r0 func(http.Handler) http.Handler
	if rf, ok := ret.Get(0).(func() func(http.Handler) http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func(http.Handler) http.Handler)
		}
	}

	return r0
}