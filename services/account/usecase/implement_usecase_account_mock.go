// Code generated by mockery v2.42.0. DO NOT EDIT.

package usecase

import (
	context "context"
	models "tinder-cloning/models"

	mock "github.com/stretchr/testify/mock"

	schema "tinder-cloning/services/account/schema"
)

// MockAccountUseCase is an autogenerated mock type for the AccountUseCase type
type MockAccountUseCase struct {
	mock.Mock
}

// GetOne provides a mock function with given fields: ctx, filter
func (_m *MockAccountUseCase) GetOne(ctx context.Context, filter schema.AccountFilter) (*models.AccountAsProfile, error) {
	ret := _m.Called(ctx, filter)

	if len(ret) == 0 {
		panic("no return value specified for GetOne")
	}

	var r0 *models.AccountAsProfile
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, schema.AccountFilter) (*models.AccountAsProfile, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, schema.AccountFilter) *models.AccountAsProfile); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.AccountAsProfile)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, schema.AccountFilter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignUp provides a mock function with given fields: ctx, payload
func (_m *MockAccountUseCase) SignUp(ctx context.Context, payload *schema.RequestRegister) error {
	ret := _m.Called(ctx, payload)

	if len(ret) == 0 {
		panic("no return value specified for SignUp")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *schema.RequestRegister) error); ok {
		r0 = rf(ctx, payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SingIn provides a mock function with given fields: ctx, payload
func (_m *MockAccountUseCase) SingIn(ctx context.Context, payload *schema.RequestLogin) (string, error) {
	ret := _m.Called(ctx, payload)

	if len(ret) == 0 {
		panic("no return value specified for SingIn")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *schema.RequestLogin) (string, error)); ok {
		return rf(ctx, payload)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *schema.RequestLogin) string); ok {
		r0 = rf(ctx, payload)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *schema.RequestLogin) error); ok {
		r1 = rf(ctx, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockAccountUseCase creates a new instance of MockAccountUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAccountUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAccountUseCase {
	mock := &MockAccountUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
