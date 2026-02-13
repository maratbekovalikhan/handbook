package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Course struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title string             `bson:"title" json:"title"`
	Level string             `bson:"level" json:"level"`
}
