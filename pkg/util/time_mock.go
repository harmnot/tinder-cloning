// Code generated by mockery v2.42.0. DO NOT EDIT.

package util

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// MockTime is an autogenerated mock type for the Time type
type MockTime struct {
	mock.Mock
}

// GenerateDateOfBirthFromString provides a mock function with given fields: s
func (_m *MockTime) GenerateDateOfBirthFromString(s string) (time.Time, error) {
	ret := _m.Called(s)

	if len(ret) == 0 {
		panic("no return value specified for GenerateDateOfBirthFromString")
	}

	var r0 time.Time
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (time.Time, error)); ok {
		return rf(s)
	}
	if rf, ok := ret.Get(0).(func(string) time.Time); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(s)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Now provides a mock function with given fields: timeGMT
func (_m *MockTime) Now(timeGMT *int) time.Time {
	ret := _m.Called(timeGMT)

	if len(ret) == 0 {
		panic("no return value specified for Now")
	}

	var r0 time.Time
	if rf, ok := ret.Get(0).(func(*int) time.Time); ok {
		r0 = rf(timeGMT)
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// NewMockTime creates a new instance of MockTime. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTime(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTime {
	mock := &MockTime{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}