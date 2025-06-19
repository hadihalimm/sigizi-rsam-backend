package request

type CreateFood struct {
	Name               string `json:"name" binding:"required" validate:"required"`
	FoodMaterialUsages []struct {
		FoodMaterialID uint    `json:"foodMaterialID" validate:"required"`
		QuantityUsed   float64 `json:"quantityUsed" validate:"required,gt=0"`
	} `json:"foodMaterialUsages" validate:"required,dive"`
}

type UpdateFood struct {
	Name               string `json:"name" binding:"required" validate:"required"`
	FoodMaterialUsages []struct {
		FoodMaterialID uint    `json:"foodMaterialID" validate:"required"`
		QuantityUsed   float64 `json:"quantityUsed" validate:"required,gt=0"`
	} `json:"foodMaterialUsages" validate:"required,dive"`
}
