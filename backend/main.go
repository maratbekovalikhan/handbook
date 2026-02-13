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
	// Загружаем .env (локально)
	_ = godotenv.Load()

	// Проверка MONGO_URI
	log.Println("MONGO_URI:", os.Getenv("MONGO_URI"))

	// Подключение к MongoDB
	config.ConnectMongo()

	// ======== Раздача фронтенда ========
	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/", fs) // Любой запрос к / будет отдавать файлы из frontend

	// ======== API ========
	http.HandleFunc("/api/register", handlers.Register)
	http.HandleFunc("/api/login", handlers.Login)
	http.Handle("/api/profile", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetProfile)))
	http.Handle("/api/progress/update", middleware.AuthMiddleware(http.HandlerFunc(handlers.UpdateProgress)))
	http.Handle("/api/progress/me", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetProgress)))

	// ======== Порт ========
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Локально по умолчанию
	}

	log.Println("Server running on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
