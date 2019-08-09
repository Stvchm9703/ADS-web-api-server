package model

import (
	"gopkg.in/mgo.v2/bson"
)

type LocationMod struct {
	Address *string `bson:"address,omitempty" json:"address,omitempty"`
}

type LocationModP LocationMod

func (s *LocationModP) SetBSON(raw bson.Raw) error {
	return raw.Unmarshal(s)
}
