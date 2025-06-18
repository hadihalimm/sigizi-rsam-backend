package request

type CreateFoodMaterial struct {
	Name         string  `json:"name" binding:"required" validate:"required"`
	Unit         string  `json:"unit" binding:"required" validate:"required"`
	PricePerUnit float64 `json:"pricePerUnit" binding:"required" validate:"required"`
}

type UpdateFoodMaterial struct {
	Name         string  `json:"name" binding:"required" validate:"required"`
	Unit         string  `json:"unit" binding:"required" validate:"required"`
	PricePerUnit float64 `json:"pricePerUnit" binding:"required" validate:"required"`
}
