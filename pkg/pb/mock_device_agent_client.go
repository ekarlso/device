// Code generated by mockery v2.9.4. DO NOT EDIT.

package pb

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// MockDeviceAgentClient is an autogenerated mock type for the DeviceAgentClient type
type MockDeviceAgentClient struct {
	mock.Mock
}

// ConfigureJITA provides a mock function with given fields: ctx, in, opts
func (_m *MockDeviceAgentClient) ConfigureJITA(ctx context.Context, in *ConfigureJITARequest, opts ...grpc.CallOption) (*ConfigureJITAResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ConfigureJITAResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ConfigureJITARequest, ...grpc.CallOption) *ConfigureJITAResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ConfigureJITAResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ConfigureJITARequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAgentConfiguration provides a mock function with given fields: ctx, in, opts
func (_m *MockDeviceAgentClient) GetAgentConfiguration(ctx context.Context, in *GetAgentConfigurationRequest, opts ...grpc.CallOption) (*GetAgentConfigurationResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *GetAgentConfigurationResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetAgentConfigurationRequest, ...grpc.CallOption) *GetAgentConfigurationResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetAgentConfigurationResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetAgentConfigurationRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, in, opts
func (_m *MockDeviceAgentClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *LoginResponse
	if rf, ok := ret.Get(0).(func(context.Context, *LoginRequest, ...grpc.CallOption) *LoginResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*LoginResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *LoginRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Logout provides a mock function with given fields: ctx, in, opts
func (_m *MockDeviceAgentClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *LogoutResponse
	if rf, ok := ret.Get(0).(func(context.Context, *LogoutRequest, ...grpc.CallOption) *LogoutResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*LogoutResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *LogoutRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetAgentConfiguration provides a mock function with given fields: ctx, in, opts
func (_m *MockDeviceAgentClient) SetAgentConfiguration(ctx context.Context, in *SetAgentConfigurationRequest, opts ...grpc.CallOption) (*SetAgentConfigurationResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *SetAgentConfigurationResponse
	if rf, ok := ret.Get(0).(func(context.Context, *SetAgentConfigurationRequest, ...grpc.CallOption) *SetAgentConfigurationResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*SetAgentConfigurationResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *SetAgentConfigurationRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Status provides a mock function with given fields: ctx, in, opts
func (_m *MockDeviceAgentClient) Status(ctx context.Context, in *AgentStatusRequest, opts ...grpc.CallOption) (DeviceAgent_StatusClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 DeviceAgent_StatusClient
	if rf, ok := ret.Get(0).(func(context.Context, *AgentStatusRequest, ...grpc.CallOption) DeviceAgent_StatusClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(DeviceAgent_StatusClient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *AgentStatusRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
