// Code generated by mockery v2.14.0. DO NOT EDIT.

package mock

import (
	context "context"

	antibrut "github.com/romsar/antibrut"
	mock "github.com/stretchr/testify/mock"
)

// RateLimiter is an autogenerated mock type for the rateLimiter type
type RateLimiter struct {
	mock.Mock
}

// Check provides a mock function with given fields: ctx, c, val
func (_m *RateLimiter) Check(ctx context.Context, c antibrut.LimitationCode, val string) error {
	ret := _m.Called(ctx, c, val)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, antibrut.LimitationCode, string) error); ok {
		r0 = rf(ctx, c, val)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Reset provides a mock function with given fields: ctx, filter
func (_m *RateLimiter) Reset(ctx context.Context, filter antibrut.BucketFilter) error {
	ret := _m.Called(ctx, filter)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, antibrut.BucketFilter) error); ok {
		r0 = rf(ctx, filter)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRateLimiter interface {
	mock.TestingT
	Cleanup(func())
}

// NewRateLimiter creates a new instance of RateLimiter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRateLimiter(t mockConstructorTestingTNewRateLimiter) *RateLimiter {
	mock := &RateLimiter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
