package models

import "time"

type Account struct {
	ID           string     `json:"id"`
	Email        string     `json:"email"`
	Gender       string     `json:"gender"`
	Username     string     `json:"username"`
	PasswordHash string     `json:"password_hash"`
	CreatedAt    time.Time  `json:"created_at"`
	IsVerified   bool       `json:"is_verified"`
	Location     *string    `json:"location"`
	Bio          *string    `json:"bio"`
	Avatar       *string    `json:"avatar"`
	DateOfBirth  *time.Time `json:"date_of_birth"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

type AccountAsProfile struct {
	ID          string     `json:"id"`
	Email       string     `json:"email"`
	Gender      string     `json:"gender"`
	Username    string     `json:"username"`
	CreatedAt   time.Time  `json:"created_at"`
	IsVerified  bool       `json:"is_verified"`
	Location    *string    `json:"location"`
	Bio         *string    `json:"bio"`
	Avatar      *string    `json:"avatar"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
