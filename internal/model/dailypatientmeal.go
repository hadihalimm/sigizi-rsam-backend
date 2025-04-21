package model

import "time"

type DailyPatientMeal struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Date       time.Time `gorm:"not null;index" json:"date"`
	PatientID  uint      `gorm:"not null;index" json:"patientID"`
	RoomID     uint      `gorm:"not null;index" json:"roomID"`
	MealTypeID uint      `gorm:"not null;index" json:"mealTypeID"`
	Notes      string    `json:"notes"`
}
