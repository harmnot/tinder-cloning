package repository

import (
	"context"
	"database/sql"
	"tinder-cloning/models"
	"tinder-cloning/services/account/schema"
)

type Repository interface {
	CreateOne(ctx context.Context, sqlTx *sql.Tx, data *models.Account) (string, *sql.Tx, error)
	FindOne(ctx context.Context, props schema.AccountFilter) (*models.Account, error)
}
