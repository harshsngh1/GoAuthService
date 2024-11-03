package routes

import (
	"GoAuthService/internals/handlers"
	"GoAuthService/internals/middleware"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/api/v1/signup", handlers.SignUpHandler).Methods("POST")
	router.HandleFunc("/api/v1/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/api/v1/welcome", middleware.AuthMiddleware(handlers.WelcomeToApplication)).Methods("GET")
	router.HandleFunc("/api/v1/revoke", handlers.RevokeTokenHandler).Methods("POST")
	router.HandleFunc("/api/v1/refresh-token", handlers.RefreshTokenHandler).Methods("POST")
	router.HandleFunc("/api/v1/get-users", handlers.GetUserHandler).Methods("GET")
	router.HandleFunc("/api/v1/get-tokens", handlers.GetTokenHandler).Methods("GET")
}
