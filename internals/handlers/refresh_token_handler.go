package handlers

import (
	"GoAuthService/internals/models"
	"GoAuthService/utils"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Threshold for refresh
const refreshThreshold = 3 * time.Minute

func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
		return
	}

	tokenString := authHeader[len("Bearer "):]

	claims, err := utils.ValidateToken(tokenString)
	if err != nil {
		log.Println("Error validating token:", err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Check if token is close to expiry
	expiryTime := time.Unix(claims.ExpiresAt, 0)
	timeRemaining := time.Until(expiryTime)

	if timeRemaining > refreshThreshold {
		log.Printf("Token is not close to expiry, remaining time :  %v, skipping refresh", timeRemaining)
		http.Error(w, "Token is not eligible for refresh yet", http.StatusBadRequest)
		return
	}

	user := models.Users{
		UserId:  claims.UserId,
		EmailId: claims.Email,
	}

	newToken, err := utils.GenerateJWT(user)
	if err != nil {
		http.Error(w, "Failed to generate new token", http.StatusInternalServerError)
		return
	}

	utils.StoreToken(newToken, user.UserId)

	// Respond with the new token
	response := map[string]string{"new_token": newToken}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
