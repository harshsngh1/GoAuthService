package storage

import (
	"GoAuthService/internals/models"
	"strings"
)

var userStore = make(map[string]models.Users)

func CreateUser(user models.Users) {
	normalizedEmail := strings.ToLower(user.EmailId)
	userStore[normalizedEmail] = user
}

func GetUser(email string) (models.Users, bool) {
	normalizedEmail := strings.ToLower(email)
	user, exists := userStore[normalizedEmail]
	return user, exists
}

func GetAllUsers() ([]models.Users, error) {
	var users []models.Users
	for _, user := range userStore {
		users = append(users, user)
	}
	return users, nil
}
