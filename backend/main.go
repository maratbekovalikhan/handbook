package main

import (
	"log"
	"net/http"
	"os"

	"handbook/config"
	"handbook/handlers"
)

func main() {

	config.Connect()

	mux := http.NewServeMux()

	mux.HandleFunc("/api/courses", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			handlers.CreateCourse(w, r)
		} else if r.Method == "GET" {
			handlers.GetCourses(w, r)
		}
	})

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
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}
