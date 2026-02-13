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

type ProgressUpdate struct {
	Course    string `json:"course"`
	Theory    bool   `json:"theory"`
	Examples  bool   `json:"examples"`
	TestScore int    `json:"testScore"`
}

func UpdateProgress(w http.ResponseWriter, r *http.Request) {
	userCollection := config.DB.Collection("users")

	userID := r.Context().Value("userID").(string)

	var update ProgressUpdate
	json.NewDecoder(r.Body).Decode(&update)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if user.Progress == nil {
		user.Progress = make(map[string]models.CourseProgress)
	}

	c := user.Progress[update.Course]
	if update.Theory {
		c.Theory = true
	}
	if update.Examples {
		c.Examples = true
	}
	if update.TestScore > c.TestScore {
		c.TestScore = update.TestScore
	}

	user.Progress[update.Course] = c

	_, err = userCollection.UpdateOne(ctx, bson.M{"_id": userID},
		bson.M{"$set": bson.M{"progress": user.Progress}})
	if err != nil {
		http.Error(w, "Failed to update progress", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user.Progress)
}

func GetProgress(w http.ResponseWriter, r *http.Request) {
	userCollection := config.DB.Collection("users")

	userID := r.Context().Value("userID").(string)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if user.Progress == nil {
		user.Progress = make(map[string]models.CourseProgress)
	}

	json.NewEncoder(w).Encode(user.Progress)
}
