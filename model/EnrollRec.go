package model

import (
	"time"
)

type EnrollMod struct {
	StudentID  string
	Year       int
	CourseID   string
	EnrollDate time.Time
}
