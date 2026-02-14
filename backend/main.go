package main

import (
	"log"
	"net/http"
	"os"

	"handbook/config"
	"handbook/handlers"

	_ "handbook/docs"

	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Handbook API
// @version 1.0
// @description This is a handbook platform backend server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

func main() {
	// ==== Load .env ====
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	config.Connect()

	mux := http.NewServeMux()

	// API courses
	mux.HandleFunc("/api/courses", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			handlers.AuthMiddleware(handlers.CreateCourse)(w, r)
		} else if r.Method == "GET" {
			handlers.GetCourses(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/course", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handlers.GetCourse(w, r)
		} else if r.Method == "DELETE" {
			handlers.AuthMiddleware(handlers.DeleteCourse)(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	// API Progress
	mux.HandleFunc("/api/enroll", handlers.AuthMiddleware(handlers.Enroll))
	mux.HandleFunc("/api/progress", handlers.AuthMiddleware(handlers.GetProgress))
	mux.HandleFunc("/api/complete_section", handlers.AuthMiddleware(handlers.CompleteSection))

	// API Rating
	mux.HandleFunc("/api/rate", handlers.AuthMiddleware(handlers.RateCourse))
	mux.HandleFunc("/api/user_rating", handlers.AuthMiddleware(handlers.GetUserRating))

	// API auth
	mux.HandleFunc("/api/register", handlers.Register)
	mux.HandleFunc("/api/login", handlers.Login)
	mux.HandleFunc("/api/profile", handlers.AuthMiddleware(handlers.Profile))

	// API Certificate
	mux.HandleFunc("/api/certificate", handlers.AuthMiddleware(handlers.GenerateCertificate))

	// Swagger
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	// Static frontend
	fs := http.FileServer(http.Dir("./frontend"))
	mux.Handle("/", fs)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, enableCORS(mux)))
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}
