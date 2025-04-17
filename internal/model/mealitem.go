package model

import "time"

type MealItem struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Date       time.Time `gorm:"not null" json:"date"`
	MealTypeID uint      `gorm:"not null" json:"mealTypeID"`
	FoodID     uint      `gorm:"not null" json:"foodID"`
	Quantity   float64   `gorm:"not null" json:"quantity"`
	MealType   *MealType `json:"mealType"`
	Food       *Food     `json:"food"`
}
