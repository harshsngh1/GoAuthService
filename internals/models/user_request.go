package models

type UserRequest struct {
	EmailId  string `json:"email"`
	Password string `json:"password"`
}
