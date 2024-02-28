package models

import "time"

type Membership struct {
	ID             int        `json:"id"`
	AccountID      string     `json:"account_id"`
	MembershipType string     `json:"membership_type"`
	StartDate      *time.Time `json:"start_date"`
	EndDate        *time.Time `json:"end_date"`
	PaymentMethod  string     `json:"payment_method"`
}
