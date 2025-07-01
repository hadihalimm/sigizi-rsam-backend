package request

type CreateFoodMaterial struct {
	Name            string  `json:"name" binding:"required" validate:"required"`
	Unit            string  `json:"unit" binding:"required" validate:"required"`
	Category        string  `json:"category" binding:"required" validate:"required,oneof=kering basah"`
	StandardPerMeal float64 `json:"standardPerMeal" binding:"required" validate:"required"`
}

type UpdateFoodMaterial struct {
	Name            string  `json:"name" binding:"required" validate:"required"`
	Unit            string  `json:"unit" binding:"required" validate:"required"`
	Category        string  `json:"category" binding:"required" validate:"required,oneof=kering basah"`
	StandardPerMeal float64 `json:"standardPerMeal" binding:"required" validate:"required"`
}
