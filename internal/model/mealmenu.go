package model

import "time"

type MealMenu struct {
	ID                 uint      `gorm:"primaryKey" json:"id"`
	Name               string    `gorm:"not null;uniqueIndex:idx_mealmenu" json:"name"`
	Day                uint      `gorm:"not null;uniqueIndex:idx_mealmenu" json:"day"`
	Time               string    `gorm:"not null;uniqueIndex:idx_mealmenu" json:"time"`
	MealTypeID         uint      `gorm:"not null;uniqueIndex:idx_mealmenu" json:"mealTypeID"`
	Foods              []Food    `gorm:"many2many:meal_foods;" json:"foods"`
	MealType           MealType  `json:"mealType"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
	MealMenuTemplateID uint      `gorm:"not null" json:"mealMenuTemplateID"`
}

type MealMenuTemplate struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Name      string     `gorm:"unique;not null" json:"name"`
	MealMenus []MealMenu `json:"mealMenus"`
}
