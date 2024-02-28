package usecase

import (
	"context"
	"tinder-cloning/models"
	"tinder-cloning/services/account/schema"
)

type AccountUseCase interface {
	SignUp(ctx context.Context, payload *schema.RequestRegister) error
	SingIn(ctx context.Context, payload *schema.RequestLogin) (string, error)
	GetOne(ctx context.Context, filter schema.AccountFilter) (*models.AccountAsProfile, error)
}
