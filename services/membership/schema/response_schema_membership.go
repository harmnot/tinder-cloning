package schema

type FeatureMembership struct {
	Name              string `json:"name"`
	QuotaSwipes       int    `json:"quota"`
	ShowWhoCanSeeMe   bool   `json:"show_control_who_can_see_me"`
	ShowVerifiedLabel bool   `json:"show_verified_label"`
}
