package usecase_test

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
	"tinder-cloning/models"
	"tinder-cloning/services/membership/repository"
	"tinder-cloning/services/membership/usecase"
)

type mocks struct {
	mock *repository.MockRepository
}

func provideMocksMembership() *mocks {
	return &mocks{
		&repository.MockRepository{},
	}
}

func provideUseCaseMembership(m *mocks) usecase.MembershipUseCase {
	return usecase.NewMembershipUseCase(m.mock)
}

func TestImplementMembershipUseCase_CreateOne(t *testing.T) {
	type expectedCreateOne struct {
		ctx     context.Context
		payload *models.Membership
		err     error
	}
	helperTest := func(ex expectedCreateOne) *mocks {
		mc := provideMocksMembership()

		mc.mock.On("CreateOne", ex.ctx, &sql.Tx{}, ex.payload).Return(&sql.Tx{}, ex.err).Once()
		return mc
	}

	tmNow := time.Now()

	tests := []struct {
		name                string
		expected            expectedCreateOne
		funcUseCaseShouldBe func(t *testing.T, err error)
	}{
		{
			name: "CreateOne success",
			expected: expectedCreateOne{
				ctx: context.Background(),
				payload: &models.Membership{
					AccountID:      mock.Anything,
					MembershipType: usecase.LevelFree,
					StartDate:      nil,
					EndDate:        nil,
					PaymentMethod:  "",
				},
				err: nil,
			},
			funcUseCaseShouldBe: func(t *testing.T, err error) {
				assert.Nil(t, err)
			},
		},
		{
			name: "CreateOne Failed",
			expected: expectedCreateOne{
				ctx: context.Background(),
				payload: &models.Membership{
					AccountID:      mock.Anything,
					MembershipType: usecase.LevelPremium,
					StartDate:      &tmNow,
					EndDate:        &tmNow,
					PaymentMethod:  "paypal",
				},
				err: assert.AnError,
			},
			funcUseCaseShouldBe: func(t *testing.T, err error) {
				assert.NotNil(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := helperTest(tt.expected)

			useCase := provideUseCaseMembership(m)
			_, err := useCase.CreateOne(tt.expected.ctx, &sql.Tx{}, tt.expected.payload)
			tt.funcUseCaseShouldBe(t, err)
		})
	}
}

func TestImplementMembershipUseCase_GetFeatureMembership(t *testing.T) {
	type expectedGetFeatureMembership struct {
		ctx       context.Context
		accountID string
		err       error
	}
	helperTest := func(ex expectedGetFeatureMembership) *mocks {
		mc := provideMocksMembership()

		mc.mock.On("GetOne", ex.ctx, ex.accountID).Return(&models.Membership{
			AccountID:      ex.accountID,
			MembershipType: usecase.LevelFree,
			StartDate:      nil,
			EndDate:        nil,
			PaymentMethod:  "",
		}, ex.err).Once()
		return mc
	}

	tests := []struct {
		name                string
		expected            expectedGetFeatureMembership
		funcUseCaseShouldBe func(t *testing.T, err error)
	}{
		{
			name: "GetFeatureMembership success",
			expected: expectedGetFeatureMembership{
				ctx:       context.Background(),
				accountID: mock.Anything,
				err:       nil,
			},
			funcUseCaseShouldBe: func(t *testing.T, err error) {
				assert.Nil(t, err)
			},
		},
		{
			name: "GetFeatureMembership Failed",
			expected: expectedGetFeatureMembership{
				ctx:       context.Background(),
				accountID: mock.Anything,
				err:       assert.AnError,
			},
			funcUseCaseShouldBe: func(t *testing.T, err error) {
				assert.NotNil(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := helperTest(tt.expected)

			useCase := provideUseCaseMembership(m)
			_, err := useCase.GetFeatureMembership(tt.expected.ctx, tt.expected.accountID)
			tt.funcUseCaseShouldBe(t, err)
		})
	}
}
