package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Rating struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID   primitive.ObjectID `bson:"user_id" json:"user_id"`
	CourseID primitive.ObjectID `bson:"course_id" json:"course_id"`
	Score    int                `bson:"score" json:"score"` // 1-5
}
