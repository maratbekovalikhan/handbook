package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// CourseProgress хранит прогресс по каждому курсу
type CourseProgress struct {
	Theory    bool `bson:"theory" json:"theory"`
	Examples  bool `bson:"examples" json:"examples"`
	TestScore int  `bson:"testScore" json:"testScore"`
}

// User хранит данные пользователя
type User struct {
	ID       primitive.ObjectID        `bson:"_id,omitempty" json:"id"`
	Name     string                    `bson:"name" json:"name"`
	Email    string                    `bson:"email" json:"email"`
	Password string                    `bson:"password" json:"-"`
	Progress map[string]CourseProgress `bson:"progress" json:"progress"`
}
