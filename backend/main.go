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

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, relying on environment variables")
	}

	log.Println("MONGO_URI:", os.Getenv("MONGO_URI"))

	config.ConnectMongo()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Handbook API is running!"))
	})

	http.HandleFunc("/api/register", handlers.Register)
	http.HandleFunc("/api/login", handlers.Login)

	http.Handle("/api/progress/update",
		middleware.AuthMiddleware(http.HandlerFunc(handlers.UpdateProgress)))
	http.Handle("/api/progress/me",
		middleware.AuthMiddleware(http.HandlerFunc(handlers.GetProgress)))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
