package usecase

import (
	"context"
	"database/sql"
	"strings"
	"tinder-cloning/models"
	"tinder-cloning/pkg/util"
	"tinder-cloning/services/membership/repository"
	"tinder-cloning/services/membership/schema"
)

type implementMembershipUseCase struct {
	repo repository.Repository
	time util.Time
}

func NewMembershipUseCase(repo repository.Repository) MembershipUseCase {
	return &implementMembershipUseCase{repo: repo, time: util.ProvideNewTimesCustom()}
}

func (i *implementMembershipUseCase) CreateOne(ctx context.Context, sqlTx *sql.Tx, data *models.Membership) (*sql.Tx, error) {
	paymentMethod := ""
	if strings.ToLower(data.MembershipType) != LevelFree && data.PaymentMethod != "" {
		paymentMethod = data.PaymentMethod
	}

	if strings.ToLower(data.MembershipType) != LevelFree {
		data.StartDate = nil
		data.EndDate = nil
	}

	payload := &models.Membership{
		AccountID:      data.AccountID,
		MembershipType: data.MembershipType,
		StartDate:      data.StartDate,
		EndDate:        data.EndDate,
		PaymentMethod:  paymentMethod,
	}

	return i.repo.CreateOne(ctx, sqlTx, payload)
}

func (i *implementMembershipUseCase) GetFeatureMembership(ctx context.Context, accountID string) (*schema.FeatureMembership, error) {
	membership, err := i.repo.GetOne(ctx, accountID)
	if err != nil {
		return nil, err
	}

	var featureMembership schema.FeatureMembership
	switch strings.ToLower(membership.MembershipType) {
	case LevelPremium:
		featureMembership = schema.FeatureMembership{
			Name:              LevelPremium,
			QuotaSwipes:       FeatureSwipeUnlimited,
			ShowWhoCanSeeMe:   false,
			ShowVerifiedLabel: true,
		}
	case LevelGold:
		featureMembership = schema.FeatureMembership{
			Name:              LevelGold,
			QuotaSwipes:       FeatureSwipeUnlimited,
			ShowWhoCanSeeMe:   true,
			ShowVerifiedLabel: true,
		}
	default:
		featureMembership = schema.FeatureMembership{
			Name:              LevelFree,
			QuotaSwipes:       FeatureSwipeBasic,
			ShowWhoCanSeeMe:   false,
			ShowVerifiedLabel: false,
		}
	}

	timeNow := i.time.Now(nil)

	// check if end date is expired
	if membership.EndDate != nil && timeNow.After(*membership.EndDate) {
		featureMembership = schema.FeatureMembership{
			Name:              LevelFree,
			QuotaSwipes:       FeatureSwipeBasic,
			ShowWhoCanSeeMe:   false,
			ShowVerifiedLabel: false,
		}

		// update membership type to free
		err := i.repo.UpdateOne(ctx, &models.Membership{
			AccountID:      membership.AccountID,
			MembershipType: LevelFree,
			StartDate:      nil,
			EndDate:        nil,
			PaymentMethod:  "",
			UpdateAt:       &timeNow,
		})

		if err != nil {
			return nil, err
		}
	}

	return &featureMembership, nil
}

func (i *implementMembershipUseCase) UpdateOne(ctx context.Context, data *schema.UpgradeMembership) error {
	timeNow := i.time.Now(nil)
	endDate := timeNow.AddDate(0, data.HowManyMonth, 0)

	return i.repo.UpdateOne(ctx, &models.Membership{
		AccountID:      data.AccountID,
		MembershipType: data.LevelName,
		StartDate:      &timeNow,
		EndDate:        &endDate,
		PaymentMethod:  "paypal",
		UpdateAt:       &timeNow,
	})
}
