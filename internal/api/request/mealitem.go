package request

type CreateMealItem struct {
	MealTypeID uint    `json:"mealTypeID" binding:"required" validate:"required"`
	FoodID     uint    `json:"foodID" binding:"required" validate:"required"`
	Quantity   float64 `json:"quantity" binding:"required" validate:"required"`
}

type UpdateMealItem struct {
	MealTypeID uint    `json:"mealTypeID" binding:"required" validate:"required"`
	FoodID     uint    `json:"foodID" binding:"required" validate:"required"`
	Quantity   float64 `json:"quantity" binding:"required" validate:"required"`
}
