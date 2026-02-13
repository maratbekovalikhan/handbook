package models

type CourseProgress struct {
	Theory    bool `bson:"theory" json:"theory"`
	Examples  bool `bson:"examples" json:"examples"`
	TestScore int  `bson:"testScore" json:"testScore"`
}
