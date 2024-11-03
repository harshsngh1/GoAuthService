package handlers

import (
	"GoAuthService/internals/models"
	"GoAuthService/internals/storage"
	"GoAuthService/utils"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {

	var signupRequest models.UserRequest

	err := json.NewDecoder(r.Body).Decode(&signupRequest)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := validateSignUpRequest(signupRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, exists := storage.GetUser(signupRequest.EmailId); exists {
		log.Printf("User already exists: %s", signupRequest.EmailId)
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signupRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	user := models.Users{
		UserId:   utils.GenerateUUID(),
		EmailId:  signupRequest.EmailId,
		Password: string(hashedPassword),
	}

	storage.CreateUser(user)

	w.WriteHeader(http.StatusCreated)
	response := map[string]string{"message": "User created successfully"}
	json.NewEncoder(w).Encode(response)
}

func validateSignUpRequest(request models.UserRequest) error {
	if request.EmailId == "" {
		return errors.New("email is required")
	}
	if request.Password == "" {
		return errors.New("password is required")
	}
	return nil
}
