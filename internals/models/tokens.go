package models

import "time"

type Token struct {
	TokenString  string    `json:"tokenString"`
	Expiry       time.Time `json:"expiry"`
	UserId       string    `json:"userId"`
	IsRevoked    bool      `json:"isRevoked"`
	IsExpired    bool      `json:"isExpired"`
	IssuedAt     time.Time `json:"issuedAt"`
	RevokeReason string    `json:"revoke_reason,omitempty"`
}
