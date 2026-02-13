package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"handbook/config"
	"handbook/models"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	var course models.Course
	json.NewDecoder(r.Body).Decode(&course)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := config.DB.Collection("courses").InsertOne(ctx, course)
	if err != nil {
		http.Error(w, "Error saving course", 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Course created"})
}

func GetCourses(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := config.DB.Collection("courses").Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Error fetching courses", 500)
		return
	}

	var courses []models.Course
	cursor.All(ctx, &courses)

	json.NewEncoder(w).Encode(courses)
}
