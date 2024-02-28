package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"tinder-cloning/models"
	"tinder-cloning/pkg/util"
	accountSchema "tinder-cloning/services/account/schema"
	accountUseCase "tinder-cloning/services/account/usecase"
	membershipUseCase "tinder-cloning/services/membership/usecase"
	"tinder-cloning/services/swipe/repository"
	"tinder-cloning/services/swipe/schema"
)

type implementSwipesUseCase struct {
	repo              repository.Repository
	membershipUseCase membershipUseCase.MembershipUseCase
	accountUseCase    accountUseCase.AccountUseCase
	time              util.Time
}

func NewSwipesUseCase(
	repo repository.Repository,
	membershipUseCase membershipUseCase.MembershipUseCase,
	accountUseCase accountUseCase.AccountUseCase,
) SwipesUseCase {
	return &implementSwipesUseCase{
		repo:              repo,
		membershipUseCase: membershipUseCase,
		accountUseCase:    accountUseCase,
		time:              util.ProvideNewTimesCustom(),
	}
}

func (i *implementSwipesUseCase) GetAllProfile(ctx context.Context, filter schema.ProfileFilter) (*[]models.AccountAsProfile, error) {
	if filter.PerPage > 30 {
		filter.PerPage = 30
	}

	if filter.Page < 1 {
		filter.Page = 1
	}

	if filter.PerPage < 1 {
		filter.PerPage = 10
	}

	// check if total swipes today is more than 10 by today
	count, err := i.repo.CountSwipes(ctx, filter.CurrentAccountID)
	if err != nil {
		return nil, err
	}

	var profiles []models.AccountAsProfile

	// get membership feature
	featureMembership, err := i.membershipUseCase.GetFeatureMembership(ctx, filter.CurrentAccountID)
	if err != nil {
		return nil, err
	}

	if !featureMembership.ShowVerifiedLabel && filter.IsVerified != nil {
		return &profiles, errors.New("only premium and gold membership can filter by verified label")
	}

	if featureMembership.QuotaSwipes != membershipUseCase.FeatureSwipeUnlimited && count >= featureMembership.QuotaSwipes {
		return &profiles, errors.New("quota swipes today is full for this account, only premium and gold membership can swipe unlimited")
	}

	// get current account to set filter by gender
	currentAccount, errGetOne := i.accountUseCase.GetOne(ctx, accountSchema.AccountFilter{
		ID: &filter.CurrentAccountID,
	})
	if errGetOne != nil {
		return nil, errGetOne
	}

	defaultFilterGender := accountUseCase.GenderFemale
	filterGender := strings.ToLower(filter.Gender)

	if strings.ToLower(currentAccount.Gender) == accountUseCase.GenderFemale {
		defaultFilterGender = accountUseCase.GenderMale
	}

	if len(filter.Gender) != 0 && (filterGender == accountUseCase.GenderMale || filterGender == accountUseCase.GenderFemale) {
		defaultFilterGender = filter.Gender
	}

	if len(filter.Gender) != 0 && filterGender == accountUseCase.GenderAll {
		defaultFilterGender = ""
	}

	filter.Gender = defaultFilterGender

	fmt.Println("filter", filter)

	// get all profile
	accounts, errView := i.repo.ViewAllProfile(ctx, filter)
	if errView != nil {
		return nil, errView
	}

	return accounts, nil
}

func (i *implementSwipesUseCase) CreateReactionSwipes(ctx context.Context, reqData schema.RequestSwipe) error {
	if reqData.AccountID == reqData.AccountIDTarget {
		return errors.New("cannot swipe yourself")
	}

	// check if total swipes today is more than 10 by today
	count, err := i.repo.CountSwipes(ctx, reqData.AccountID)
	if err != nil {
		return err
	}

	// get membership feature
	featureMembership, errGet := i.membershipUseCase.GetFeatureMembership(ctx, reqData.AccountID)
	if errGet != nil {
		return errGet
	}

	// check if swiped today
	swiped, errCheck := i.repo.CheckSwipesToday(ctx, reqData.AccountID, reqData.AccountIDTarget)
	if errCheck != nil {
		return errCheck
	}

	if swiped {
		return errors.New("target user already swiped today")
	}

	if featureMembership.QuotaSwipes != membershipUseCase.FeatureSwipeUnlimited && count >= featureMembership.QuotaSwipes {
		return errors.New("quota swipes today is full for this account, only premium and gold membership can swipe unlimited")
	}

	payload := &models.Swipe{
		SwiperID:  reqData.AccountID,
		SwipedID:  reqData.AccountIDTarget,
		SwipeType: reqData.SwipesType,
		SwipeDate: i.time.Now(nil),
	}

	// create swipes
	err = i.repo.CreateReactionSwipes(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}
