package main

import (
	"log"
	"net/http"
	"os"

	"handbook/config"
	"handbook/handlers"

	"github.com/joho/godotenv"
)

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
