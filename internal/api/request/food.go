package request

type CreateFood struct {
	Name         string  `json:"name" binding:"required" validate:"required"`
	Unit         string  `json:"unit" binding:"required" validate:"required"`
	PricePerUnit float64 `json:"pricePerUnit" binding:"required" validate:"required"`
}

type UpdateFood struct {
	Name         string  `json:"name" binding:"required" validate:"required"`
	Unit         string  `json:"unit" binding:"required" validate:"required"`
	PricePerUnit float64 `json:"pricePerUnit" binding:"required" validate:"required"`
}
