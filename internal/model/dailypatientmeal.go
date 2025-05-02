package model

import "time"

type DailyPatientMeal struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	PatientID  uint      `gorm:"not null;index;uniqueIndex:idx_patient_date" json:"patientID"`
	RoomID     uint      `gorm:"not null;index" json:"roomID"`
	MealTypeID uint      `gorm:"not null;index" json:"mealTypeID"`
	Date       time.Time `gorm:"not null;index;uniqueIndex:idx_patient_date" json:"date"`
	Notes      string    `json:"notes"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Patient    Patient   `json:"patient"`
	Room       Room      `json:"room"`
	MealType   MealType  `json:"mealType"`
	Diets      []Diet    `gorm:"many2many:daily_patient_meal_diets;" json:"diets"`
}
