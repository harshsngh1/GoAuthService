package models

type RevokeTokenRequest struct {
	UserID string `json:"user_id"`
	Token  string `json:"token"`
	Reason string `json:"reason,omitempty"`
}
