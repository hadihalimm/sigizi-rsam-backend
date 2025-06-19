package model

import "time"

type MealMenu struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"not null" json:"name"`
	Day        string    `gorm:"not null" json:"day"`
	Time       string    `gorm:"not null" json:"time"`
	Foods      []Food    `gorm:"many2many:meal_foods;" json:"foods"`
	MealTypeID uint      `json:"mealTypeID"`
	MealType   MealType  `json:"mealType"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
