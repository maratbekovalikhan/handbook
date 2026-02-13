package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"handbook/config"
	"handbook/models"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {

	var input models.User
	json.NewDecoder(r.Body).Decode(&input)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var existing models.User
	err := config.DB.Collection("users").
		FindOne(ctx, bson.M{"email": input.Email}).
		Decode(&existing)

	if err == nil {
		http.Error(w, "Email already exists", 400)
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	input.Password = string(hash)

	_, err = config.DB.Collection("users").InsertOne(ctx, input)
	if err != nil {
		http.Error(w, "Error creating user", 500)
		return
	}

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created"})
}

func Login(w http.ResponseWriter, r *http.Request) {

	var input models.User
	json.NewDecoder(r.Body).Decode(&input)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := config.DB.Collection("users").
		FindOne(ctx, bson.M{"email": input.Email}).
		Decode(&user)

	if err != nil {
		http.Error(w, "Invalid credentials", 401)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", 401)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID.Hex(),
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func Profile(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value("userID").(string)
	objID, _ := primitive.ObjectIDFromHex(userID)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := config.DB.Collection("users").
		FindOne(ctx, bson.M{"_id": objID}).
		Decode(&user)

	if err != nil {
		http.Error(w, "User not found", 404)
		return
	}

	user.Password = ""
	json.NewEncoder(w).Encode(user)
}
