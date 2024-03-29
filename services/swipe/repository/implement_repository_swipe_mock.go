// Code generated by mockery v2.42.0. DO NOT EDIT.

package repository

import (
	context "context"
	models "tinder-cloning/models"

	mock "github.com/stretchr/testify/mock"

	schema "tinder-cloning/services/swipe/schema"
)

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

// CheckSwipesToday provides a mock function with given fields: ctx, accountID, accountIDTarget
func (_m *MockRepository) CheckSwipesToday(ctx context.Context, accountID string, accountIDTarget string) (bool, error) {
	ret := _m.Called(ctx, accountID, accountIDTarget)

	if len(ret) == 0 {
		panic("no return value specified for CheckSwipesToday")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (bool, error)); ok {
		return rf(ctx, accountID, accountIDTarget)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, accountID, accountIDTarget)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, accountID, accountIDTarget)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountSwipes provides a mock function with given fields: ctx, accountID
func (_m *MockRepository) CountSwipes(ctx context.Context, accountID string) (int, error) {
	ret := _m.Called(ctx, accountID)

	if len(ret) == 0 {
		panic("no return value specified for CountSwipes")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (int, error)); ok {
		return rf(ctx, accountID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) int); ok {
		r0 = rf(ctx, accountID)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, accountID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateReactionSwipes provides a mock function with given fields: ctx, payload
func (_m *MockRepository) CreateReactionSwipes(ctx context.Context, payload *models.Swipe) error {
	ret := _m.Called(ctx, payload)

	if len(ret) == 0 {
		panic("no return value specified for CreateReactionSwipes")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Swipe) error); ok {
		r0 = rf(ctx, payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ViewAllProfile provides a mock function with given fields: ctx, props
func (_m *MockRepository) ViewAllProfile(ctx context.Context, props schema.ProfileFilter) (*[]models.AccountAsProfile, error) {
	ret := _m.Called(ctx, props)

	if len(ret) == 0 {
		panic("no return value specified for ViewAllProfile")
	}

	var r0 *[]models.AccountAsProfile
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, schema.ProfileFilter) (*[]models.AccountAsProfile, error)); ok {
		return rf(ctx, props)
	}
	if rf, ok := ret.Get(0).(func(context.Context, schema.ProfileFilter) *[]models.AccountAsProfile); ok {
		r0 = rf(ctx, props)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]models.AccountAsProfile)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, schema.ProfileFilter) error); ok {
		r1 = rf(ctx, props)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
