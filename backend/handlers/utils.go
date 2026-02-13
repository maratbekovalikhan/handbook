package handlers

import (
	"handbook/config"

	"go.mongodb.org/mongo-driver/mongo"
)

// Геттер коллекции пользователей
func getUserCollection() *mongo.Collection {
	return config.DB.Collection("users")
}
