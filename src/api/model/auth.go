package models

type AuthResponse struct {
	Token      string `json:"token"`
	UserType   string `json:"user_type"`
	UserTypeID int    `json:"user_type_id,omitempty"`
	User       User   `json:"user"`
}
