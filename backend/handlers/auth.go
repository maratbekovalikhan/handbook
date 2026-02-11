package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"handbook/config"
	"handbook/models"
	"handbook/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	hashed, _ := utils.HashPassword(user.Password)
	user.Password = hashed

	_, err := config.DB.Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		http.Error(w, "User exists", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var input models.User
	var user models.User

	json.NewDecoder(r.Body).Decode(&input)

	err := config.DB.Collection("users").
		FindOne(context.Background(), bson.M{"email": input.Email}).
		Decode(&user)

	if err != nil || !utils.CheckPassword(user.Password, input.Password) {
		http.Error(w, "Invalid credentials", 401)
		return
	}

	token, _ := utils.GenerateJWT(user.ID)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
