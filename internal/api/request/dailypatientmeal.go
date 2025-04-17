package request

import "time"

type CreateDailyPatientMeal struct {
	Date       time.Time `json:"date" binding:"required" validate:"required"`
	PatientID  uint      `json:"patientID" binding:"required" validate:"required"`
	RoomID     uint      `json:"roomID" binding:"required" validate:"required"`
	MealTypeID uint      `json:"mealTypeID" binding:"required" validate:"required"`
	Notes      string    `json:"notes"`
}

type UpdateDailyPatientMeal struct {
	Date       time.Time `json:"date" binding:"required" validate:"required"`
	PatientID  uint      `json:"patientID" binding:"required" validate:"required"`
	RoomID     uint      `json:"roomID" binding:"required" validate:"required"`
	MealTypeID uint      `json:"mealTypeID" binding:"required" validate:"required"`
	Notes      string    `json:"notes"`
}
