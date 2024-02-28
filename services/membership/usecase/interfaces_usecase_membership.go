package usecase

import (
	"context"
	"database/sql"
	"tinder-cloning/models"
	"tinder-cloning/services/membership/schema"
)

type MembershipUseCase interface {
	CreateOne(ctx context.Context, sqlTx *sql.Tx, data *models.Membership) (*sql.Tx, error)
	GetFeatureMembership(ctx context.Context, accountID string) (*schema.FeatureMembership, error)
}
