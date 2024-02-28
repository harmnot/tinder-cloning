package schema

type UpgradeMembership struct {
	AccountID    string `json:"account_id,omitempty"`
	LevelName    string `json:"level_name"`
	HowManyMonth int    `json:"how_many_month"`
}
