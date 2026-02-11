package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"handbook/config"
	"handbook/models"
	"handbook/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateProgress(w http.ResponseWriter, r *http.Request) {
	var input models.Progress
	json.NewDecoder(r.Body).Decode(&input)

	userID, _ := primitive.ObjectIDFromHex(r.Context().Value("userId").(string))
	input.UserID = userID
	input.Percent = utils.CalculateProgress(input.Theory, input.Examples, input.TestScore)

	filter := bson.M{"userId": userID, "course": input.Course}
	update := bson.M{"$set": input}

	config.DB.Collection("progress").UpdateOne(context.Background(), filter, update)
	w.WriteHeader(http.StatusOK)
}

func GetProgress(w http.ResponseWriter, r *http.Request) {
	userID, _ := primitive.ObjectIDFromHex(r.Context().Value("userId").(string))

	cursor, _ := config.DB.Collection("progress").
		Find(context.Background(), bson.M{"userId": userID})

	var result []models.Progress
	cursor.All(context.Background(), &result)
	json.NewEncoder(w).Encode(result)
}
