package handlers

import (
	"GoAuthService/internals/storage"
	"encoding/json"
	"log"
	"net/http"
)

// Handler to get all users from storage
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	users, err := storage.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		log.Println("Error retrieving users:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
