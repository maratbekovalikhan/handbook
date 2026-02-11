package main

import (
	"log"
	"net/http"

	"handbook/config"
	"handbook/handlers"
	"handbook/middleware"
)

func main() {
	config.ConnectMongo()

	http.HandleFunc("/api/register", handlers.Register)
	http.HandleFunc("/api/login", handlers.Login)

	http.Handle("/api/progress/update",
		middleware.AuthMiddleware(http.HandlerFunc(handlers.UpdateProgress)))
	http.Handle("/api/progress/me",
		middleware.AuthMiddleware(http.HandlerFunc(handlers.GetProgress)))

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
