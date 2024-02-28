// Code generated by mockery v2.42.0. DO NOT EDIT.

package usecase

import (
	context "context"
	models "tinder-cloning/models"

	mock "github.com/stretchr/testify/mock"

	schema "tinder-cloning/services/swipe/schema"
)

// MockSwipesUseCase is an autogenerated mock type for the SwipesUseCase type
type MockSwipesUseCase struct {
	mock.Mock
}

// CreateReactionSwipes provides a mock function with given fields: ctx, payload
func (_m *MockSwipesUseCase) CreateReactionSwipes(ctx context.Context, payload schema.RequestSwipe) error {
	ret := _m.Called(ctx, payload)

	if len(ret) == 0 {
		panic("no return value specified for CreateReactionSwipes")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, schema.RequestSwipe) error); ok {
		r0 = rf(ctx, payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllProfile provides a mock function with given fields: ctx, filter
func (_m *MockSwipesUseCase) GetAllProfile(ctx context.Context, filter schema.ProfileFilter) (*[]models.AccountAsProfile, error) {
	ret := _m.Called(ctx, filter)

	if len(ret) == 0 {
		panic("no return value specified for GetAllProfile")
	}

	var r0 *[]models.AccountAsProfile
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, schema.ProfileFilter) (*[]models.AccountAsProfile, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, schema.ProfileFilter) *[]models.AccountAsProfile); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]models.AccountAsProfile)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, schema.ProfileFilter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockSwipesUseCase creates a new instance of MockSwipesUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSwipesUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSwipesUseCase {
	mock := &MockSwipesUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
