package model

import "time"

type DailyPatientMeal struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Date       time.Time `gorm:"not null" json:"date"`
	PatientID  uint      `gorm:"not null" json:"patientID"`
	RoomID     uint      `gorm:"not null" json:"roomID"`
	MealTypeID uint      `gorm:"not null" json:"mealTypeID"`
	Notes      string    `json:"notes"`
}
