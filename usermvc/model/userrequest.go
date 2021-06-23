package model

type UserResquest struct {
	Email        string `json:"email"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Entitlements string `json:"entitlements"`
}

type UserResponse struct {
	Status  int
	Message string
}
