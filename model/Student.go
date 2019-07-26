package model

import (
	"time"
)

type StudentMod struct {
	StudentID string
	StuName   string
	DOB       time.Time
}
