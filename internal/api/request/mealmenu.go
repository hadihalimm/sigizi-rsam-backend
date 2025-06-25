package request

import "time"

type CreateMealMenu struct {
	Name               string `json:"name" binding:"required" validate:"required"`
	Day                uint   `json:"day" binding:"required" validate:"required"`
	Time               string `json:"time" binding:"required" validate:"required,oneof=pagi siang sore"`
	MealTypeID         uint   `json:"mealTypeID" binding:"required" validate:"required"`
	MealMenuTemplateID uint   `json:"mealMenuTemplateID" binding:"required" validate:"required"`
	FoodIDs            []uint `json:"foodIDs" validate:"required,min=1,dive,gt=0"`
}

type UpdateMealMenu struct {
	Name               string `json:"name" binding:"required" validate:"required"`
	Day                uint   `json:"day" binding:"required" validate:"required"`
	Time               string `json:"time" binding:"required" validate:"required,oneof=pagi siang sore"`
	MealTypeID         uint   `json:"mealTypeID" binding:"required" validate:"required"`
	MealMenuTemplateID uint   `json:"mealMenuTemplateID" binding:"required" validate:"required"`
	FoodIDs            []uint `json:"foodIDs" validate:"required,min=1,dive,gt=0"`
}

type CreateMealMenuTemplate struct {
	Name string `json:"name" binding:"required" validate:"required"`
}

type UpdateMealMenuTemplate struct {
	Name string `json:"name" binding:"required" validate:"required"`
}

type CreateMenuTemplateSchedule struct {
	Date               time.Time `json:"date" binding:"required" validate:"required"`
	MealMenuTemplateID uint      `json:"mealMenuTemplateID" binding:"required" validate:"required"`
}

type UpdateMenuTemplateSchedule struct {
	MealMenuTemplateID uint `json:"mealMenuTemplateID" binding:"required" validate:"required"`
}
