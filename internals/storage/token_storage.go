package storage

import (
	"GoAuthService/internals/models"
	"errors"
	"sync"
	"time"
)

var tokenStore = make(map[string]models.Token)
var mu sync.Mutex

// StoreToken stores the newly generated tokens
func StoreToken(userId string, token models.Token) {
	mu.Lock()
	defer mu.Unlock()
	tokenStore[userId] = token
}

// GetTokenByUserId fetches token on the basis of userId
func GetTokenByUserId(userId string) (models.Token, bool) {
	mu.Lock()
	defer mu.Unlock()
	token, exists := tokenStore[userId]
	return token, exists
}

func GetAllTokens() ([]models.Token, error) {
	var tokens []models.Token
	for _, token := range tokenStore {
		tokens = append(tokens, token)
	}
	return tokens, nil
}

func IsTokenExpired(token models.Token) bool {
	return time.Now().After(token.Expiry)
}

func RevokeToken(userID, token, reason string) error {
	mu.Lock()
	defer mu.Unlock()
	storedToken, exists := tokenStore[userID]
	if !exists || storedToken.TokenString != token {
		return errors.New("token not found or does not match")
	}

	storedToken.IsRevoked = true
	storedToken.RevokeReason = reason
	tokenStore[userID] = storedToken

	return nil
}

func IsTokenRevoked(userID, token string) bool {
	storedToken, exists := tokenStore[userID]
	if !exists || storedToken.TokenString != token {
		return false
	}
	return storedToken.IsRevoked
}

func GetRevokedToken(userID string) (models.Token, bool) {
	mu.Lock()
	defer mu.Unlock()
	for _, token := range tokenStore {
		if token.UserId == userID && token.IsRevoked {
			return token, true
		}
	}
	return models.Token{}, false
}
