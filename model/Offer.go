package model

type OfferMod struct {
	_id             string
	Department      DepartmentMod
	Course          CoursesMod
	Year            int
	ClassSize       int
	AvailablePlaces int
}
