package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"handbook/config"
	"handbook/handlers"
	"handbook/middleware"

	"github.com/joho/godotenv"
)

func main() {
	// Загружаем .env
	_ = godotenv.Load()

	// Подключаем MongoDB
	config.ConnectMongo()

	log.Println("MongoDB connected")

	// ======== Статика фронтенда ========
	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/", fs)

	// ======== API ========
	http.HandleFunc("/api/register", handlers.Register)
	http.HandleFunc("/api/login", handlers.Login)
	http.Handle("/api/profile",
		middleware.AuthMiddleware(http.HandlerFunc(handlers.GetProfile)))
	http.Handle("/api/progress/update",
		middleware.AuthMiddleware(http.HandlerFunc(handlers.UpdateProgress)))
	http.Handle("/api/progress/me",
		middleware.AuthMiddleware(http.HandlerFunc(handlers.GetProgress)))

	// Используем порт из Render или по умолчанию 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on :" + port)

	srv := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
