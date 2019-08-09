package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type StudentMod struct {
	ID        *bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	StudentID *string        `bson:"student_id,omitempty"json:"student_id,omitempty"`
	StuName   *string        `bson:"student_name,omitempty" json:"student_name,omitempty"`
	DOB       *time.Time     `bson:"dob,omitempty" json:"dob,omitempty"`
	CreatedAt *time.Time     `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt *time.Time     `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
