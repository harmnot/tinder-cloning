package usecase_test

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
	"tinder-cloning/models"
	"tinder-cloning/services/account/repository"
	"tinder-cloning/services/account/schema"
	"tinder-cloning/services/account/usecase"
	mockMembershipUseCase "tinder-cloning/services/membership/usecase"
)

type mocks struct {
	mock                  *repository.MockRepository
	mockMembershipUseCase *mockMembershipUseCase.MockMembershipUseCase
}

func provideMocksAccount() *mocks {
	return &mocks{
		&repository.MockRepository{},
		&mockMembershipUseCase.MockMembershipUseCase{},
	}
}

func provideUseCaseAccount(m *mocks) usecase.AccountUseCase {
	return usecase.NewAccountUseCase(m.mock, m.mockMembershipUseCase)
}

func TestImplementAccountUseCase_SignUp(t *testing.T) {
	type expectedSignUp struct {
		errFind             error
		errCreate           error
		errCreateMembership error
		id                  *string
		payloadMember       *models.Membership
		payload             *schema.RequestRegister
		ctx                 context.Context
		account             *models.Account
	}

	helperTest := func(ex expectedSignUp) *mocks {
		mc := provideMocksAccount()

		mc.mock.On("FindOne", ex.ctx, schema.AccountFilter{
			Email:    ex.payload.Email,
			Username: ex.payload.Username,
			ID:       nil,
		}).Return(ex.account, ex.errFind).Once()
		mc.mock.On("CreateOne", ex.ctx, nil, ex.payload).Return("uuid_dummy", &sql.Tx{}, ex.errCreate).Once()
		mc.mockMembershipUseCase.On("CreateOne", ex.ctx, &sql.Tx{}, ex.payloadMember).Return(&sql.Tx{}, ex.errCreateMembership).Once()
		return mc
	}

	tests := []struct {
		name                string
		expected            expectedSignUp
		funcUseCaseShouldBe func(t *testing.T, payload *schema.RequestRegister, err error)
	}{
		{
			name: "Success",
			expected: expectedSignUp{
				errFind:   nil,
				errCreate: nil,
				payload: &schema.RequestRegister{
					Gender: "MALE",
					RequestLogin: schema.RequestLogin{
						Email:    "email@mail.com",
						Password: "password",
					},
					Username: "username_123",
				},
				account: &models.Account{},
				ctx:     context.Background(),
			},
			funcUseCaseShouldBe: func(t *testing.T, payload *schema.RequestRegister, err error) {
				assert.Error(t, err)
			},
		},
		{
			name: "Error Validation",
			expected: expectedSignUp{
				errFind:   nil,
				errCreate: nil,
				payload: &schema.RequestRegister{
					Gender: "MALE",
					RequestLogin: schema.RequestLogin{
						Email:    "email@gmail.com",
						Password: "password1233",
					},
					Username: "username_wr$ng",
				},
				ctx: context.Background(),
			},
			funcUseCaseShouldBe: func(t *testing.T, payload *schema.RequestRegister, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name: "Error From Database",
			expected: expectedSignUp{
				errFind:   nil,
				errCreate: assert.AnError,
				payload: &schema.RequestRegister{
					RequestLogin: schema.RequestLogin{
						Email:    "my@gmail.io",
						Password: "password1233@jbf",
					},
					Username: "username_valid2355",
					Gender:   "MALE",
					Location: "",
					Bio:      "iam here",
					Avatar:   "",
				},
				account: &models.Account{},
				ctx:     context.Background(),
			},
			funcUseCaseShouldBe: func(t *testing.T, payload *schema.RequestRegister, err error) {
				assert.NotNil(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocks := helperTest(tt.expected)
			useCase := provideUseCaseAccount(mocks)
			tt.funcUseCaseShouldBe(t, tt.expected.payload, useCase.SignUp(tt.expected.ctx, tt.expected.payload))
		})
	}
}

func TestImplementAccountUseCase_SignIn(t *testing.T) {
	type expectedSignIn struct {
		errFind error
		id      *string
		payload *schema.RequestLogin
		ctx     context.Context
		account *models.Account
	}

	helperTest := func(ex expectedSignIn) *mocks {
		mc := provideMocksAccount()

		mc.mock.On("FindOne", ex.ctx, schema.AccountFilter{
			Email: ex.payload.Email,
			ID:    nil,
		}).Return(ex.account, ex.errFind).Once()
		return mc
	}

	tests := []struct {
		name                string
		expected            expectedSignIn
		funcUseCaseShouldBe func(t *testing.T, payload *schema.RequestLogin, err error)
	}{
		{
			name: "Success",
			expected: expectedSignIn{
				errFind: nil,
				payload: &schema.RequestLogin{
					Email:    "email@mail.com",
					Password: "password",
				},
				account: &models.Account{},
				ctx:     context.Background(),
			},
			funcUseCaseShouldBe: func(t *testing.T, payload *schema.RequestLogin, err error) {
				assert.Error(t, err)
			},
		},
		{
			name: "Error From Database",
			expected: expectedSignIn{
				errFind: assert.AnError,
				payload: &schema.RequestLogin{
					Email:    "email@mail.com",
					Password: "password",
				},
				account: &models.Account{},
				ctx:     context.Background(),
			},
			funcUseCaseShouldBe: func(t *testing.T, payload *schema.RequestLogin, err error) {
				assert.NotNil(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocks := helperTest(tt.expected)
			useCase := provideUseCaseAccount(mocks)
			_, err := useCase.SingIn(tt.expected.ctx, tt.expected.payload)
			tt.funcUseCaseShouldBe(t, tt.expected.payload, err)
		})
	}
}
