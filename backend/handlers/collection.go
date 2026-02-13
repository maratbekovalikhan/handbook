package handlers

import (
	"handbook/config"

	"go.mongodb.org/mongo-driver/mongo"
)

// Универсальная функция для получения коллекции пользователей
func getUserCollection() *mongo.Collection {
	return config.DB.Collection("users")
}
