package utils

import (
	"GoAuthService/internals/models"
	"GoAuthService/internals/storage"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const secret = "secret"
const defaultExpiry = 6 * time.Minute

func GenerateUUID() string {
	return uuid.NewString()
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT(user models.Users) (string, error) {

	claims := models.CustomClaims{
		Email:  user.EmailId,
		UserId: user.UserId,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(defaultExpiry).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) (models.CustomClaims, error) {
	var claims models.CustomClaims

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return claims, errors.New("invalid token")
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return claims, errors.New("token expired")
	}

	return claims, nil
}

func StoreToken(token string, userId string) {
	tokenData := models.Token{
		TokenString: token,
		Expiry:      time.Now().Add(defaultExpiry),
		UserId:      userId,
		IsRevoked:   false,
		IssuedAt:    time.Now(),
	}
	storage.StoreToken(userId, tokenData)
}
