package request

import "time"

type CreateMealItem struct {
	Date       time.Time `json:"date" binding:"required" validate:"required"`
	MealTypeID uint      `json:"mealTypeID" binding:"required" validate:"required"`
	FoodID     uint      `json:"foodID" binding:"required" validate:"required"`
	Quantity   float64   `json:"quantity" binding:"required" validate:"required"`
}

type UpdateMealItem struct {
	Date       time.Time `json:"date" binding:"required" validate:"required"`
	MealTypeID uint      `json:"mealTypeID" binding:"required" validate:"required"`
	FoodID     uint      `json:"foodID" binding:"required" validate:"required"`
	Quantity   float64   `json:"quantity" binding:"required" validate:"required"`
}
