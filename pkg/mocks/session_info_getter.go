// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	pb "github.com/nais/device/pkg/pb"
	mock "github.com/stretchr/testify/mock"
)

// SessionInfoGetter is an autogenerated mock type for the SessionInfoGetter type
type SessionInfoGetter struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0, _a1
func (_m *SessionInfoGetter) Execute(_a0 context.Context, _a1 string) (*pb.Session, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.Session
	if rf, ok := ret.Get(0).(func(context.Context, string) *pb.Session); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Session)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
