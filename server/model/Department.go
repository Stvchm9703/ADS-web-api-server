package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type DepartmentMod struct {
	ID        *bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	DeptID    *string        `bson:"dept_id,omitempty" json:"dept_id,omitempty"`
	DeptName  *string        `bson:"dept_name,omitempty" json:"dept_name,omitempty"`
	Location  *LocationModP  `bson:"location,omitempty" json:"location,omitempty"`
	CreatedAt *time.Time     `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt *time.Time     `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
