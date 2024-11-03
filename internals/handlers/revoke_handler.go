package handlers

import (
	"GoAuthService/internals/models"
	"GoAuthService/internals/storage"
	"encoding/json"
	"log"
	"net/http"
)

func RevokeTokenHandler(w http.ResponseWriter, r *http.Request) {

	var revokeRequest models.RevokeTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&revokeRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := storage.RevokeToken(revokeRequest.UserID, revokeRequest.Token, revokeRequest.Reason)
	if err != nil {
		log.Printf("Failed to revoke token for user %s: %s", revokeRequest.UserID, err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	log.Printf("Token revoked successfully for user %s", revokeRequest.UserID)
	response := map[string]string{
		"status":  "success",
		"message": "Token revoked successfully",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
