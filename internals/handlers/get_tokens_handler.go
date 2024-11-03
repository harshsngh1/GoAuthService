package handlers

import (
	"GoAuthService/internals/storage"
	"encoding/json"
	"log"
	"net/http"
)

// Handler to get all tokens from storage
func GetTokenHandler(w http.ResponseWriter, r *http.Request) {
	tokens, err := storage.GetAllTokens()
	if err != nil {
		http.Error(w, "Failed to retrieve tokens", http.StatusInternalServerError)
		log.Println("Error retrieving tokens:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tokens)
}
