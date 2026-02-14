package handlers

import (
	"context"
	"encoding/json"
	"math"
	"net/http"
	"time"

	"handbook/config"
	"handbook/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func RateCourse(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	userObjID, _ := primitive.ObjectIDFromHex(userID)

	var input struct {
		CourseID string `json:"course_id"`
		Score    int    `json:"score"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", 400)
		return
	}

	if input.Score < 1 || input.Score > 5 {
		http.Error(w, "Score must be between 1 and 5", 400)
		return
	}

	courseObjID, _ := primitive.ObjectIDFromHex(input.CourseID)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 1. Update or Insert the user's rating
	filter := bson.M{"user_id": userObjID, "course_id": courseObjID}
	update := bson.M{"$set": bson.M{"score": input.Score}}
	opts := options.Update().SetUpsert(true)

	_, err := config.DB.Collection("ratings").UpdateOne(ctx, filter, update, opts)
	if err != nil {
		http.Error(w, "Error saving rating", 500)
		return
	}

	// 2. Recalculate Average for the Course
	// Fetch all ratings for this course
	cursor, err := config.DB.Collection("ratings").Find(ctx, bson.M{"course_id": courseObjID})
	if err != nil {
		http.Error(w, "Error calculating average", 500)
		return
	}

	var ratings []models.Rating
	if err = cursor.All(ctx, &ratings); err != nil {
		http.Error(w, "Error processing ratings", 500)
		return
	}

	var totalScore int
	for _, r := range ratings {
		totalScore += r.Score
	}

	count := len(ratings)
	var avg float64
	if count > 0 {
		avg = float64(totalScore) / float64(count)
		// Round to 1 decimal place
		avg = math.Round(avg*10) / 10
	}

	// 3. Update Course document
	_, err = config.DB.Collection("courses").UpdateOne(ctx, bson.M{"_id": courseObjID}, bson.M{
		"$set": bson.M{
			"average_rating": avg,
			"rating_count":   count,
		},
	})

	if err != nil {
		http.Error(w, "Error updating course stats", 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":        "Rating saved",
		"average_rating": avg,
		"rating_count":   count,
	})
}

func GetUserRating(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	userObjID, _ := primitive.ObjectIDFromHex(userID)
	
	courseIDStr := r.URL.Query().Get("course_id")
	courseObjID, _ := primitive.ObjectIDFromHex(courseIDStr)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var rating models.Rating
	err := config.DB.Collection("ratings").FindOne(ctx, bson.M{
		"user_id":   userObjID,
		"course_id": courseObjID,
	}).Decode(&rating)

	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"score": 0}) // Not rated yet
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"score": rating.Score})
}
