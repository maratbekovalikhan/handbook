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
	_ = godotenv.Load() // загружаем .env

	log.Println("MONGO_URI:", os.Getenv("MONGO_URI"))

	config.ConnectMongo()

	// ======== Статика фронтенда ========
	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/", fs)

	// ======== API ========
	http.HandleFunc("/api/register", handlers.Register)
	http.HandleFunc("/api/login", handlers.Login)
	http.Handle("/api/progress/update",
		middleware.AuthMiddleware(http.HandlerFunc(handlers.UpdateProgress)))
	http.Handle("/api/progress/me",
		middleware.AuthMiddleware(http.HandlerFunc(handlers.GetProgress)))
	http.Handle("/api/profile",
		middleware.AuthMiddleware(http.HandlerFunc(handlers.GetProfile)))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
