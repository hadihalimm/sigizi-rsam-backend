package model

import "time"

type MealItem struct {
	// Date       time.Time `gorm:"not null" json:"date"`
	ID         uint      `gorm:"primaryKey" json:"id"`
	MealTypeID uint      `gorm:"not null" json:"mealTypeID"`
	FoodID     uint      `gorm:"not null" json:"foodID"`
	Quantity   float64   `gorm:"not null" json:"quantity"`
	MealType   *MealType `json:"mealType"`
	Food       *Food     `json:"food"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
