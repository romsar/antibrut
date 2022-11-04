// Code generated by mockery v2.14.0. DO NOT EDIT.

package mock

import (
	context "context"

	antibrut "github.com/romsar/antibrut"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the service type
type Service struct {
	mock.Mock
}

// AddIPToBlackList provides a mock function with given fields: ctx, subnet
func (_m *Service) AddIPToBlackList(ctx context.Context, subnet antibrut.Subnet) error {
	ret := _m.Called(ctx, subnet)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, antibrut.Subnet) error); ok {
		r0 = rf(ctx, subnet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddIPToWhiteList provides a mock function with given fields: ctx, subnet
func (_m *Service) AddIPToWhiteList(ctx context.Context, subnet antibrut.Subnet) error {
	ret := _m.Called(ctx, subnet)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, antibrut.Subnet) error); ok {
		r0 = rf(ctx, subnet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Check provides a mock function with given fields: ctx, login, pass, ip
func (_m *Service) Check(ctx context.Context, login antibrut.Login, pass antibrut.Password, ip antibrut.IP) error {
	ret := _m.Called(ctx, login, pass, ip)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, antibrut.Login, antibrut.Password, antibrut.IP) error); ok {
		r0 = rf(ctx, login, pass, ip)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteIPFromBlackList provides a mock function with given fields: ctx, subnet
func (_m *Service) DeleteIPFromBlackList(ctx context.Context, subnet antibrut.Subnet) error {
	ret := _m.Called(ctx, subnet)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, antibrut.Subnet) error); ok {
		r0 = rf(ctx, subnet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteIPFromWhiteList provides a mock function with given fields: ctx, subnet
func (_m *Service) DeleteIPFromWhiteList(ctx context.Context, subnet antibrut.Subnet) error {
	ret := _m.Called(ctx, subnet)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, antibrut.Subnet) error); ok {
		r0 = rf(ctx, subnet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Reset provides a mock function with given fields: ctx, login, ip
func (_m *Service) Reset(ctx context.Context, login antibrut.Login, ip antibrut.IP) error {
	ret := _m.Called(ctx, login, ip)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, antibrut.Login, antibrut.IP) error); ok {
		r0 = rf(ctx, login, ip)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewService interface {
	mock.TestingT
	Cleanup(func())
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t mockConstructorTestingTNewService) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
