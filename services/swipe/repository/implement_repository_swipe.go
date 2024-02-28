package repository

import (
	"context"
	"database/sql"
	"fmt"
	"tinder-cloning/models"
	"tinder-cloning/services/swipe/schema"
)

type swipeRepositoryImpl struct {
	db *sql.DB
}

func NewSwipesRepositoryImpl(db *sql.DB) Repository {
	return &swipeRepositoryImpl{db: db}
}

func (s *swipeRepositoryImpl) CountSwipes(ctx context.Context, accountID string) (int, error) {
	query := `
		SELECT COUNT(*)
		FROM swipes
		WHERE swiper_id = $1 AND DATE(swipe_date) = CURRENT_DATE;
	`
	var count int
	err := s.db.QueryRowContext(ctx, query, accountID).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("querying swipes: %v", err)
	}
	return count, nil
}

func (s *swipeRepositoryImpl) CheckSwipesToday(ctx context.Context, accountID string, accountIDTarget string) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1
			FROM swipes
			WHERE swiper_id = $1 AND swiped_id = $2 AND DATE(swipe_date) = CURRENT_DATE
		);
	`
	var exists bool
	err := s.db.QueryRowContext(ctx, query, accountID, accountIDTarget).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("querying swipes: %v", err)
	}
	return exists, nil
}

func (s *swipeRepositoryImpl) ViewAllProfile(ctx context.Context, props schema.ProfileFilter) (*[]models.AccountAsProfile, error) {
	query := `SELECT ac.id, ac.username, ac.email, ac.avatar, ac.bio, ac.date_of_birth,
            ac.gender, ac.location, ac.is_verified, ac.created_at, ac.updated_at
            FROM accounts AS ac
            WHERE ac.id NOT IN (
                SELECT swiped_id
                FROM swipes
                WHERE swiper_id = $1 AND DATE(swipe_date) = CURRENT_DATE
            ) AND ac.id != $1`

	queryArgs := []interface{}{props.CurrentAccountID}

	if props.IsVerified != nil && props.Gender == "" {
		query += " AND ac.is_verified = $2 LIMIT $3 OFFSET ($4 - 1) * $3"
		queryArgs = append(queryArgs, *props.IsVerified, props.PerPage, props.Page)
	} else if props.IsVerified != nil && props.Gender != "" {
		query += " AND ac.is_verified = $2 AND lower(ac.gender) = $3 LIMIT $4 OFFSET ($5 - 1) * $4"
		queryArgs = append(queryArgs, *props.IsVerified, props.Gender, props.PerPage, props.Page)
	} else if props.IsVerified == nil && props.Gender != "" {
		query += " AND lower(ac.gender) = $2 LIMIT $3 OFFSET ($4 - 1) * $3"
		queryArgs = append(queryArgs, props.Gender, props.PerPage, props.Page)
	} else {
		query += " LIMIT $2 OFFSET ($3 - 1) * $2"
		queryArgs = append(queryArgs, props.PerPage, props.Page)
	}

	rows, err := s.db.QueryContext(ctx, query, queryArgs...)
	if err != nil {
		return nil, fmt.Errorf("querying accounts to display: %v", err)
	}

	return s.scanAccounts(rows)
}

func (s *swipeRepositoryImpl) scanAccounts(rows *sql.Rows) (*[]models.AccountAsProfile, error) {
	var accounts []models.AccountAsProfile
	for rows.Next() {
		var account models.AccountAsProfile
		err := rows.Scan(
			&account.ID, &account.Username, &account.Email, &account.Avatar,
			&account.Bio, &account.DateOfBirth, &account.Gender, &account.Location,
			&account.IsVerified, &account.CreatedAt, &account.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scanning account rows: %v", err)
		}
		accounts = append(accounts, account)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("handling account rows: %v", err)
	}

	return &accounts, nil
}

func (s *swipeRepositoryImpl) CreateReactionSwipes(ctx context.Context, payload *models.Swipe) error {
	query := `
		INSERT INTO swipes(swiper_id, swiped_id, swipe_type, swipe_date)
		VALUES($1, $2, $3, $4)
	`
	_, err := s.db.ExecContext(ctx, query, payload.SwiperID, payload.SwipedID, payload.SwipeType, payload.SwipeDate)
	if err != nil {
		return fmt.Errorf("inserting swipe reaction: %v", err)
	}
	return nil
}
