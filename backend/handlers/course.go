package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"handbook/config"
	"handbook/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	var course models.Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		http.Error(w, "Invalid input JSON", 400)
		return
	}

	if course.Sections == nil {
		course.Sections = []models.Section{}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get user from context (set by AuthMiddleware)
	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusInternalServerError)
		return
	}

	var user models.User
	err = config.DB.Collection("users").FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	course.AuthorID = objID
	course.AuthorName = user.Name

	_, err = config.DB.Collection("courses").InsertOne(ctx, course)
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

func GetCourse(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id parameter", 400)
		return
	}

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, "Invalid course ID", 400)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var course models.Course
	err = config.DB.Collection("courses").FindOne(ctx, bson.M{"_id": id}).Decode(&course)
	if err != nil {
		http.Error(w, "Course not found", 404)
		return
	}

	json.NewEncoder(w).Encode(course)
}
