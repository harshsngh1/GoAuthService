package handlers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	UserId  string `json:"userId"`
}

func WelcomeToApplication(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("userId")

	if userId == "" {
		http.Error(w, "userId header is missing", http.StatusBadRequest)
		return
	}

	response := Response{
		Message: "Hey user, Welcome to the Application!",
		UserId:  userId,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
