package models

type CreateSessionRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateAccountRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Session struct {
	Token string `json:"token"`
}

type Account struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}
