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

	// ======== Статика фронтенда ========
	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/", fs)

	// ======== API ========
	http.HandleFunc("/api/register", handlers.Register)
	http.HandleFunc("/api/login", handlers.Login)
	http.Handle("/api/profile", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetProfile)))
	http.Handle("/api/progress/update", middleware.AuthMiddleware(http.HandlerFunc(handlers.UpdateProgress)))
	http.Handle("/api/progress/me", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetProgress)))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
