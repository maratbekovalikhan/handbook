package main

import (
	"log"
	"net/http"
	"os"

	"handbook/config"
	"handbook/handlers"
	"handbook/middleware"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	config.ConnectMongo()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Handbook API is running!"))
	})

	http.HandleFunc("/api/register", handlers.Register)
	http.HandleFunc("/api/login", handlers.Login)
	http.Handle("/api/profile", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetProfile)))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
