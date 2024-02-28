package repository

import (
	"context"
	"database/sql"
	"tinder-cloning/models"
)

type membershipRepositoryImpl struct {
	db *sql.DB
}

func NewMembershipRepositoryImpl(db *sql.DB) Repository {
	return &membershipRepositoryImpl{db: db}
}

func (m *membershipRepositoryImpl) CreateOne(ctx context.Context, sqlTx *sql.Tx, data *models.Membership) (*sql.Tx, error) {
	var tx *sql.Tx
	var err error
	if sqlTx == nil {
		tx, err = m.db.BeginTx(ctx, nil)
		if err != nil {
			return nil, err
		}
	} else {
		tx = sqlTx
	}

	_, err = tx.ExecContext(ctx, `
		INSERT INTO memberships (account_id, membership_type, start_date, end_date, payment_method) 
		VALUES ($1, $2, $3, $4, $5)`, data.AccountID, data.MembershipType, data.StartDate, data.EndDate, data.PaymentMethod,
	)

	if err != nil {
		errRl := tx.Rollback()
		if errRl != nil {
			return nil, errRl
		}
		return nil, err
	}

	return tx, nil
}

func (m *membershipRepositoryImpl) GetOne(ctx context.Context, accountID string) (*models.Membership, error) {
	var membership models.Membership
	err := m.db.QueryRowContext(ctx, `
		SELECT account_id, membership_type, start_date, end_date, payment_method
		FROM memberships
		WHERE account_id = $1`, accountID,
	).Scan(&membership.AccountID, &membership.MembershipType, &membership.StartDate, &membership.EndDate, &membership.PaymentMethod)

	if err != nil {
		return nil, err
	}

	return &membership, nil
}
