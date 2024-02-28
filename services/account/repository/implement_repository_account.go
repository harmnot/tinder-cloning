package repository

import (
	"context"
	"database/sql"
	"strings"
	"tinder-cloning/models"
	"tinder-cloning/services/account/schema"
)

type accountRepositoryImpl struct {
	db *sql.DB
}

func NewAccountRepositoryImpl(db *sql.DB) Repository {
	return &accountRepositoryImpl{db: db}
}

func (a *accountRepositoryImpl) CreateOne(ctx context.Context, sqlTx *sql.Tx, data *models.Account) (string, *sql.Tx, error) {
	var tx *sql.Tx
	var err error
	if sqlTx == nil {
		tx, err = a.db.BeginTx(ctx, nil)
		if err != nil {
			return "", nil, err
		}
	} else {
		tx = sqlTx
	}

	var lastInsertId string

	err = tx.QueryRowContext(ctx, `
	INSERT INTO accounts (email, username, password_hash, is_verified, location, bio, gender, avatar, date_of_birth, created_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`,
		data.Email, data.Username, data.PasswordHash, data.IsVerified, data.Location, data.Bio, data.Gender, data.Avatar, data.DateOfBirth, data.CreatedAt,
	).Scan(&lastInsertId)

	if err != nil {
		errRl := tx.Rollback()
		if errRl != nil {
			return "", nil, errRl
		}
		return "", nil, err
	}

	return lastInsertId, tx, nil
}

func (a *accountRepositoryImpl) FindOne(ctx context.Context, props schema.AccountFilter) (*models.Account, error) {
	row := a.db.QueryRowContext(ctx, "SELECT id, email, username, password_hash, gender, bio, avatar, date_of_birth, location, created_at, updated_at FROM accounts WHERE lower(username) = $1 OR lower(email) = $2", strings.ToLower(props.Username), strings.ToLower(props.Email))
	if props.ID != nil {
		row = a.db.QueryRowContext(ctx, "SELECT id, email, username, password_hash, gender, bio, avatar, date_of_birth, location, created_at, updated_at FROM accounts WHERE lower(username) = $1 OR lower(email) = $2 OR id = $3", strings.ToLower(props.Username), strings.ToLower(props.Email), *props.ID)
	}
	var account models.Account
	err := row.Scan(&account.ID, &account.Email, &account.Username, &account.PasswordHash, &account.Gender, &account.Bio, &account.Avatar, &account.DateOfBirth, &account.Location, &account.CreatedAt, &account.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if account.ID == "" {
		return nil, nil
	}
	return &account, nil
}
