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
	// Загружаем .env
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using environment variables")
	}

	// Логируем URI для проверки
	log.Println("MONGO_URI:", os.Getenv("MONGO_URI"))

	// Подключаем MongoDB
	config.ConnectMongo()

	// ======== Раздача фронтенда ========
	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/", http.StripPrefix("/", fs)) // важно для корректной обработки вложенных путей

	// ======== API ========
	http.HandleFunc("/api/register", handlers.Register)
	http.HandleFunc("/api/login", handlers.Login)
	http.Handle("/api/profile", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetProfile)))
	http.Handle("/api/progress/update", middleware.AuthMiddleware(http.HandlerFunc(handlers.UpdateProgress)))
	http.Handle("/api/progress/me", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetProgress)))

	// ======== Порт ========
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // локально
	}

	log.Println("Server running on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
