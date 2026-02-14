package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Progress struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID              primitive.ObjectID `bson:"user_id" json:"user_id"`
	CourseID            primitive.ObjectID `bson:"course_id" json:"course_id"`
	CompletedSectionIDs []string           `bson:"completed_section_ids" json:"completed_section_ids"`
	IsFinished          bool               `bson:"is_finished" json:"is_finished"`
}
