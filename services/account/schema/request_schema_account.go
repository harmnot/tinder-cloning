package schema

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestRegister struct {
	RequestLogin
	Username    string `json:"username"`
	Gender      string `json:"gender"`
	Location    string `json:"location"`
	Bio         string `json:"bio"`
	Avatar      string `json:"avatar"`
	DateOfBirth string `json:"date_of_birth"`
}

type AccountFilter struct {
	Email    string
	Username string
	ID       *string
}
