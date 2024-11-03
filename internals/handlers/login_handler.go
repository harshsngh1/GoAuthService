package handlers

import (
	"GoAuthService/internals/models"
	"GoAuthService/internals/storage"
	"GoAuthService/utils"
	"encoding/json"
	"log"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginRequest models.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		log.Printf("Error decoding login request: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, exists := storage.GetUser(loginRequest.EmailId)
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if !utils.CheckPasswordHash(loginRequest.Password, user.Password) {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	utils.StoreToken(token, user.UserId)

	response := map[string]interface{}{
		"message": "Login successful",
		"token":   token,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
