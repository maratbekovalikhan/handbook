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

// GetProgress returns the user's progress for a specific course
func GetProgress(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	userObjID, _ := primitive.ObjectIDFromHex(userID)
	
	courseIDStr := r.URL.Query().Get("course_id")
	if courseIDStr == "" {
		http.Error(w, "Missing course_id", 400)
		return
	}
	courseObjID, _ := primitive.ObjectIDFromHex(courseIDStr)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var progress models.Progress
	err := config.DB.Collection("progress").FindOne(ctx, bson.M{
		"user_id":   userObjID,
		"course_id": courseObjID,
	}).Decode(&progress)

	if err != nil {
		// If not found, return empty object (not started yet)
		json.NewEncoder(w).Encode(map[string]interface{}{"started": false})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"started":               true,
		"completed_section_ids": progress.CompletedSectionIDs,
		"is_finished":           progress.IsFinished,
	})
}

// Enroll starts the course for the user
func Enroll(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	userObjID, _ := primitive.ObjectIDFromHex(userID)

	var input struct {
		CourseID string `json:"course_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", 400)
		return
	}
	courseObjID, _ := primitive.ObjectIDFromHex(input.CourseID)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if already enrolled
	count, _ := config.DB.Collection("progress").CountDocuments(ctx, bson.M{
		"user_id":   userObjID,
		"course_id": courseObjID,
	})

	if count > 0 {
		json.NewEncoder(w).Encode(map[string]string{"message": "Already enrolled"})
		return
	}

	newProgress := models.Progress{
		UserID:              userObjID,
		CourseID:            courseObjID,
		CompletedSectionIDs: []string{},
		IsFinished:          false,
	}

	_, err := config.DB.Collection("progress").InsertOne(ctx, newProgress)
	if err != nil {
		http.Error(w, "Error enrolling", 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Enrolled successfully"})
}

// CompleteSection marks a section as done
func CompleteSection(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	userObjID, _ := primitive.ObjectIDFromHex(userID)

	var input struct {
		CourseID  string `json:"course_id"`
		SectionID string `json:"section_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", 400)
		return
	}
	courseObjID, _ := primitive.ObjectIDFromHex(input.CourseID)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 1. Update Progress
	update := bson.M{
		"$addToSet": bson.M{"completed_section_ids": input.SectionID},
	}
	
	_, err := config.DB.Collection("progress").UpdateOne(ctx, bson.M{
		"user_id":   userObjID,
		"course_id": courseObjID,
	}, update)

	if err != nil {
		http.Error(w, "Error updating progress", 500)
		return
	}

	// 2. Check if Course is Finished
	// We need to fetch the course to know total sections
	var course models.Course
	err = config.DB.Collection("courses").FindOne(ctx, bson.M{"_id": courseObjID}).Decode(&course)
	
	if err == nil {
		// Fetch updated progress
		var progress models.Progress
		config.DB.Collection("progress").FindOne(ctx, bson.M{
			"user_id":   userObjID,
			"course_id": courseObjID,
		}).Decode(&progress)

		if len(progress.CompletedSectionIDs) >= len(course.Sections) {
			config.DB.Collection("progress").UpdateOne(ctx, bson.M{
				"user_id":   userObjID,
				"course_id": courseObjID,
			}, bson.M{"$set": bson.M{"is_finished": true}})
		}
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Section completed"})
}
