package usecase_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"tinder-cloning/models"
	"tinder-cloning/pkg/util"
	accountSchema "tinder-cloning/services/account/schema"
	mockAccountUseCase "tinder-cloning/services/account/usecase"
	membershipSchema "tinder-cloning/services/membership/schema"
	mockMembershipUseCase "tinder-cloning/services/membership/usecase"
	"tinder-cloning/services/swipe/repository"
	"tinder-cloning/services/swipe/schema"
	"tinder-cloning/services/swipe/usecase"
)

type mocks struct {
	mock                  *repository.MockRepository
	mockMembershipService *mockMembershipUseCase.MockMembershipUseCase
	mockAccountService    *mockAccountUseCase.MockAccountUseCase
	mockTime              *util.MockTime
}

func provideMocksSwipe() *mocks {
	return &mocks{
		mock:                  &repository.MockRepository{},
		mockMembershipService: &mockMembershipUseCase.MockMembershipUseCase{},
		mockAccountService:    &mockAccountUseCase.MockAccountUseCase{},
		mockTime:              &util.MockTime{},
	}
}

func provideUseCaseSwipe(m *mocks) usecase.SwipesUseCase {
	return usecase.NewSwipesUseCase(m.mock, m.mockMembershipService, m.mockAccountService, m.mockTime)
}

func TestImplementSwipesUseCase_CreateReactionSwipes(t *testing.T) {
	type expectedCreateReactionSwipes struct {
		ctx                           context.Context
		payload                       schema.RequestSwipe
		responseFeature               *membershipSchema.FeatureMembership
		resultCountSwipe              int
		isSwipedAlready               bool
		err, errGetCount              error
		errGetFeature, errCheckSwiped error
	}
	helperTest := func(ex expectedCreateReactionSwipes) *mocks {
		mc := provideMocksSwipe()

		mc.mock.On("CountSwipes", ex.ctx, ex.payload.AccountID).Return(ex.resultCountSwipe, ex.errGetCount).Once()

		mc.mockMembershipService.On("GetFeatureMembership", ex.ctx, ex.payload.AccountID).Return(ex.responseFeature, ex.errGetFeature).Once()
		mc.mock.On("CheckSwipesToday", ex.ctx, ex.payload.AccountID, ex.payload.AccountIDTarget).Return(ex.isSwipedAlready, ex.errCheckSwiped).Once()

		timeNow := time.Now().UTC()
		var timeGmt *int
		mc.mockTime.On("Now", timeGmt).Return(timeNow).Once()

		mc.mock.On("CreateReactionSwipes", ex.ctx, &models.Swipe{
			SwiperID:  ex.payload.AccountID,
			SwipedID:  ex.payload.AccountIDTarget,
			SwipeType: ex.payload.SwipesType,
			SwipeDate: timeNow,
		}).Return(ex.err).Once()
		return mc
	}

	tests := []struct {
		name                string
		expected            expectedCreateReactionSwipes
		funcUseCaseShouldBe func(t *testing.T, err error)
	}{
		{
			name: "success create reaction swipes",
			expected: expectedCreateReactionSwipes{
				ctx: context.Background(),
				payload: schema.RequestSwipe{
					AccountID:       "uuid1233",
					AccountIDTarget: "uuid9999999",
					SwipesType:      "like",
				},
				responseFeature: &membershipSchema.FeatureMembership{
					Name:              mockMembershipUseCase.LevelFree,
					QuotaSwipes:       10,
					ShowWhoCanSeeMe:   false,
					ShowVerifiedLabel: false,
				},
				resultCountSwipe: 9,
				isSwipedAlready:  false,
				err:              nil,
				errGetCount:      nil,
				errGetFeature:    nil,
				errCheckSwiped:   nil,
			},
			funcUseCaseShouldBe: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
		{
			name: "cannot swipe yourself",
			expected: expectedCreateReactionSwipes{
				ctx: context.Background(),
				payload: schema.RequestSwipe{
					AccountID:       "uuid1233",
					AccountIDTarget: "uuid1233",
					SwipesType:      "pass",
				},
			},
			funcUseCaseShouldBe: func(t *testing.T, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name: "target already swiped today",
			expected: expectedCreateReactionSwipes{
				ctx: context.Background(),
				payload: schema.RequestSwipe{
					AccountID:       "uuid1233",
					AccountIDTarget: "uuid9999999",
					SwipesType:      "like",
				},
				responseFeature: &membershipSchema.FeatureMembership{
					Name:              mockMembershipUseCase.LevelFree,
					QuotaSwipes:       10,
					ShowWhoCanSeeMe:   false,
					ShowVerifiedLabel: false,
				},
				resultCountSwipe: 3,
				isSwipedAlready:  true,
				err:              nil,
				errGetCount:      nil,
				errGetFeature:    nil,
				errCheckSwiped:   nil,
			},
			funcUseCaseShouldBe: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
		{
			name: "user limit swipes today",
			expected: expectedCreateReactionSwipes{
				ctx: context.Background(),
				payload: schema.RequestSwipe{
					AccountID:       "uuid1233",
					AccountIDTarget: "uuid9999999",
					SwipesType:      "like",
				},
				responseFeature: &membershipSchema.FeatureMembership{
					Name:              mockMembershipUseCase.LevelFree,
					QuotaSwipes:       10,
					ShowWhoCanSeeMe:   false,
					ShowVerifiedLabel: false,
				},
				resultCountSwipe: 10,
				isSwipedAlready:  true,
				err:              nil,
				errGetCount:      nil,
				errGetFeature:    nil,
				errCheckSwiped:   nil,
			},
			funcUseCaseShouldBe: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
		{
			name: "get error when create reaction swipes from database",
			expected: expectedCreateReactionSwipes{
				ctx: context.Background(),
				payload: schema.RequestSwipe{
					AccountID:       "uuid1233",
					AccountIDTarget: "uuid9999999",
					SwipesType:      "pass",
				},
				responseFeature: &membershipSchema.FeatureMembership{
					Name:              mockMembershipUseCase.LevelFree,
					QuotaSwipes:       10,
					ShowWhoCanSeeMe:   false,
					ShowVerifiedLabel: false,
				},
				resultCountSwipe: 9,
				isSwipedAlready:  true,
				err:              assert.AnError,
				errGetCount:      nil,
				errGetFeature:    nil,
				errCheckSwiped:   nil,
			},
			funcUseCaseShouldBe: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := helperTest(tt.expected)

			uc := provideUseCaseSwipe(mc)
			err := uc.CreateReactionSwipes(tt.expected.ctx, tt.expected.payload)

			tt.funcUseCaseShouldBe(t, err)
		})
	}
}

func TestImplementSwipesUseCase_GetAllProfile(t *testing.T) {
	type expectedGetAllProfile struct {
		ctx                context.Context
		filter             schema.ProfileFilter
		responseFeature    *membershipSchema.FeatureMembership
		resultCountSwipe   int
		errCountSwipes     error
		err, errGetAccount error
		errGetFeature      error
	}
	helperTest := func(ex expectedGetAllProfile) *mocks {
		mc := provideMocksSwipe()

		mc.mock.On("CountSwipes", ex.ctx, ex.filter.CurrentAccountID).Return(ex.resultCountSwipe, ex.errCountSwipes).Once()
		mc.mockMembershipService.On("GetFeatureMembership", ex.ctx, ex.filter.CurrentAccountID).Return(ex.responseFeature, ex.errGetFeature).Once()
		mc.mockAccountService.On("GetOne", ex.ctx, accountSchema.AccountFilter{
			ID: &ex.filter.CurrentAccountID,
		}).Return(&models.AccountAsProfile{}, ex.errGetAccount).Once()
		mc.mock.On("ViewAllProfile", ex.ctx, ex.filter).Return(nil, ex.err).Once()
		return mc
	}

	tests := []struct {
		name                string
		expected            expectedGetAllProfile
		funcUseCaseShouldBe func(t *testing.T, err error)
	}{
		{
			name: "success get all profile",
			expected: expectedGetAllProfile{
				ctx: context.Background(),
				filter: schema.ProfileFilter{
					CurrentAccountID: "uuid1233",
					PerPage:          10,
					Page:             1,
					IsVerified:       nil,
					Gender:           "female",
				},
				responseFeature: &membershipSchema.FeatureMembership{
					Name:              mockMembershipUseCase.LevelFree,
					QuotaSwipes:       10,
					ShowWhoCanSeeMe:   false,
					ShowVerifiedLabel: false,
				},
				resultCountSwipe: 0,
				errCountSwipes:   nil,
				err:              nil,
				errGetAccount:    nil,
				errGetFeature:    nil,
			},
			funcUseCaseShouldBe: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
		{
			name: "empty data because limit swipes today",
			expected: expectedGetAllProfile{
				ctx: context.Background(),
				filter: schema.ProfileFilter{
					CurrentAccountID: "uuid1233",
					PerPage:          10,
					Page:             1,
					IsVerified:       nil,
					Gender:           "male",
				},
				responseFeature: &membershipSchema.FeatureMembership{
					Name:              mockMembershipUseCase.LevelFree,
					QuotaSwipes:       10,
					ShowWhoCanSeeMe:   false,
					ShowVerifiedLabel: false,
				},
				resultCountSwipe: 10,
				errCountSwipes:   nil,
				err:              nil,
				errGetAccount:    nil,
				errGetFeature:    nil,
			},
			funcUseCaseShouldBe: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := helperTest(tt.expected)

			uc := provideUseCaseSwipe(mc)
			_, err := uc.GetAllProfile(tt.expected.ctx, tt.expected.filter)

			tt.funcUseCaseShouldBe(t, err)
		})
	}

}
