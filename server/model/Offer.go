package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type OfferMod struct {
	ID              *bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Department      *DepartmentMod `bson:"department,omitempty"json:"department,omitempty"`
	Course          *CourseMod     `bson:"course,omitempty" json:"course,omitempty"`
	Year            *int           `bson:"year,omitempty" json:"year,omitempty"`
	ClassSize       *int           `bson:"class_size,omitempty" json:"class_size,omitempty"`
	AvailablePlaces *int           `bson:"available_places,omitempty"json:"available_places,omitempty"`
	CreatedAt       *time.Time     `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt       *time.Time     `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
