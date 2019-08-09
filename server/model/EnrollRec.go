package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type EnrollMod struct {
	ID         *bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	StudentID  *string        `bson:"student_id,omitempty" json:"student_id,omitempty"`
	Year       *int           `bson:"year,omitempty" json:"year,omitempty"`
	CourseID   *string        `bson:"course_id,omitempty" json:"course_id,omitempty"`
	EnrollDate *time.Time     `bson:"enroll_date,omitempty" json:"enroll_date,omitempty"`
	CreatedAt  *time.Time     `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt  *time.Time     `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
