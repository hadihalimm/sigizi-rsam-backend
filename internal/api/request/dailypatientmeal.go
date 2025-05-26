package request

import "time"

type CreateDailyPatientMeal struct {
	PatientID       uint      `json:"patientID" binding:"required" validate:"required"`
	RoomID          uint      `json:"roomID" binding:"required" validate:"required"`
	MealTypeID      uint      `json:"mealTypeID" binding:"required" validate:"required"`
	Date            time.Time `json:"date" binding:"required" validate:"required"`
	Notes           string    `json:"notes"`
	DietIDs         []uint    `json:"dietIDs"`
	IsNewlyAdmitted bool      `json:"isNewlyAdmitted"`
}

type UpdateDailyPatientMeal struct {
	PatientID       uint   `json:"patientID" binding:"required" validate:"required"`
	RoomID          uint   `json:"roomID" binding:"required" validate:"required"`
	MealTypeID      uint   `json:"mealTypeID" binding:"required" validate:"required"`
	Notes           string `json:"notes"`
	DietIDs         []uint `json:"dietIDs"`
	IsNewlyAdmitted bool   `json:"isNewlyAdmitted"`
}
