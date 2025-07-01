package model

import "time"

type MealMenu struct {
	ID                 uint      `gorm:"primaryKey" json:"id"`
	Name               string    `gorm:"not null;" json:"name"`
	Day                uint      `gorm:"not null;" json:"day"`
	Time               string    `gorm:"not null;" json:"time"`
	MenuType           string    `gorm:"not null;" json:"menu_type"`
	MealTypeID         uint      `gorm:"not null;" json:"mealTypeID"`
	Foods              []Food    `gorm:"many2many:meal_foods;" json:"foods"`
	SnackID            *uint     `gorm:"" json:"snackID"`
	Snack              *Snack    `gorm:"" json:"snack"`
	MealMenuTemplateID uint      `gorm:"not null;" json:"mealMenuTemplateID"`
	MealType           MealType  `json:"mealType"`
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

type MenuTemplateSchedule struct {
	ID                 uint             `gorm:"primaryKey" json:"id"`
	Date               time.Time        `gorm:"uniqueIndex:idx_template_date" json:"date"`
	MealMenuTemplateID uint             `gorm:"uniqueIndex:idx_template_date" json:"mealMenuTemplateID"`
	MealMenuTemplate   MealMenuTemplate `json:"mealMenuTemplate"`
}
