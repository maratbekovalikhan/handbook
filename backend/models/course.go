package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Course struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Level       string             `bson:"level" json:"level"`
	Description string             `bson:"description" json:"description"`
	PhotoURL    string             `bson:"photo_url" json:"photo_url"`
	GeneralInfo string             `bson:"general_info" json:"general_info"`
	AuthorID    primitive.ObjectID `bson:"author_id" json:"author_id"`
	AuthorName  string             `bson:"author_name" json:"author_name"`
}
