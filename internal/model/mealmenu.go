package model

import "time"

type MealMenu struct {
	ID                 uint      `gorm:"primaryKey" json:"id"`
	Name               string    `gorm:"not null;" json:"name"`
	Day                uint      `gorm:"not null;" json:"day"`
	Time               string    `gorm:"not null;" json:"time"`
	MealTypeID         uint      `gorm:"not null;" json:"mealTypeID"`
	Foods              []Food    `gorm:"many2many:meal_foods;" json:"foods"`
	MealType           MealType  `json:"mealType"`
	MealMenuTemplateID uint      `gorm:"not null;" json:"mealMenuTemplateID"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
}

type MealMenuTemplate struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Name      string     `gorm:"unique;not null" json:"name"`
	MealMenus []MealMenu `json:"mealMenus"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}
