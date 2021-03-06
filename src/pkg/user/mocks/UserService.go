// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// GetInfo provides a mock function with given fields:
func (_m *UserService) GetInfo() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetUserInfo provides a mock function with given fields:
func (_m *UserService) GetUserInfo() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Login provides a mock function with given fields:
func (_m *UserService) Login() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
