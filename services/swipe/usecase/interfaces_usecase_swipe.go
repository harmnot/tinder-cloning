package usecase

import (
	"context"
	"tinder-cloning/models"
	"tinder-cloning/services/swipe/schema"
)

type SwipesUseCase interface {
	GetAllProfile(ctx context.Context, filter schema.ProfileFilter) (*[]models.AccountAsProfile, error)
	CreateReactionSwipes(ctx context.Context, payload schema.RequestSwipe) error
}
