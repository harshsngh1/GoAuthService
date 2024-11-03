package main

import (
	"GoAuthService/internals/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Starting server on Port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
