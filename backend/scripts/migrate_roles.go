package main

import (
	"context"
	"log"
	"time"

	"handbook/config"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	// 1. Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 2. Connect to DB
	config.Connect()
	db := config.DB
	usersCollection := db.Collection("users")

	// 3. Define the update logic
	// We update only those users who don't have the "role" field yet
	filter := bson.M{"role": bson.M{"$exists": false}}
	update := bson.M{"$set": bson.M{"role": "user"}}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result, err := usersCollection.UpdateMany(ctx, filter, update)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Printf("Migration completed successfully!")
	log.Printf("Matched %v documents and updated %v documents.", result.MatchedCount, result.ModifiedCount)
}
