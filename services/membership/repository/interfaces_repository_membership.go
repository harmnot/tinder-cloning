package repository

import (
	"context"
	"database/sql"
	"tinder-cloning/models"
)

type Repository interface {
	CreateOne(ctx context.Context, sqlTx *sql.Tx, data *models.Membership) (*sql.Tx, error)
	GetOne(ctx context.Context, accountID string) (*models.Membership, error)
}
