package model

type UserResquest struct {
Email        string `json:"email"`
FirstName    string `json:"firstName"`
LastName     string `json:"lastName"`
Entitlements string `json:"entitlements,omitempty"`
}

type UserResponse struct {
	Status int
	Message string
}
