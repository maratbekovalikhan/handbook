package main

import (
	"log"
	"net/http"
)

func main() {

	logChan := make(chan string, 10)

	go logger(logChan)

	repo := NewArticleRepository()
	service := NewArticleService(repo, logChan)
	handler := NewHandler(service)

	http.HandleFunc("/articles", handler.Articles)
	http.HandleFunc("/articles/", handler.Article)

	log.Println("Server started :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func logger(ch chan string) {

	for msg := range ch {
		log.Println("[LOG]", msg)
	}
}
