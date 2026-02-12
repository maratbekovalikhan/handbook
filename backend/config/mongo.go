package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("mongodb+srv://maratbekovalikhan_db_user:11qwertyuiop@cluster0.g21xhch.mongodb.net/handbook?retryWrites=true&w=majority")))
	if err != nil {
		log.Fatal(err)
	}

	DB = client.Database("handbook")
	log.Println("MongoDB connected")
}
