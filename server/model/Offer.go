package model

type OfferMod struct {
	_id             string
	Department      DepartmentMod
	Course          CourseMod
	Year            int
	ClassSize       int
	AvailablePlaces int
}
