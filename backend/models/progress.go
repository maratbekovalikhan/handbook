package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Progress struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"userId"`
	Course    string             `bson:"course"`
	Theory    bool               `bson:"theory"`
	Examples  bool               `bson:"examples"`
	TestScore int                `bson:"testScore"`
	Percent   int                `bson:"percent"`
}
