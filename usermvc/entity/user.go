package entity

type User struct {
	Email        string `json:"email"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Entitlements string `json:"entitlements,omitempty"`
}

