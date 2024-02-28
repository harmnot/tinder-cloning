package repository

import (
	"context"
	"tinder-cloning/models"
	"tinder-cloning/services/swipe/schema"
)

type Repository interface {
	CountSwipes(ctx context.Context, accountID string) (int, error)
	CheckSwipesToday(ctx context.Context, accountID string, accountIDTarget string) (bool, error)
	ViewAllProfile(ctx context.Context, props schema.ProfileFilter) (*[]models.AccountAsProfile, error)
	CreateReactionSwipes(ctx context.Context, payload *models.Swipe) error
}
