package request

type CreateFood struct {
	MealTypeID uint    `json:"mealTypeID" binding:"required" validate:"required"`
	FoodID     uint    `json:"foodID" binding:"required" validate:"required"`
	Quantity   float64 `json:"quantity" binding:"required" validate:"required"`
}

type UpdateFood struct {
	MealTypeID uint    `json:"mealTypeID" binding:"required" validate:"required"`
	FoodID     uint    `json:"foodID" binding:"required" validate:"required"`
	Quantity   float64 `json:"quantity" binding:"required" validate:"required"`
}
