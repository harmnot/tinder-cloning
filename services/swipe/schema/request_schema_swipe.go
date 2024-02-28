package schema

type RequestSwipe struct {
	AccountID       string `json:"account_id"`
	AccountIDTarget string `json:"account_id_target"`
	SwipesType      string `json:"swipes_type"`
}

type ProfileFilter struct {
	CurrentAccountID string
	PerPage          int
	Page             int
	IsVerified       *bool
	Gender           string
}
