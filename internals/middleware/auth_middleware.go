package middleware

import (
	"GoAuthService/internals/storage"
	"GoAuthService/utils"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
			return
		}

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if storage.IsTokenRevoked(claims.UserId, tokenString) {
			http.Error(w, "Token has been revoked", http.StatusUnauthorized)
			return
		}

		r.Header.Set("userId", claims.UserId)

		next.ServeHTTP(w, r)
	}
}
