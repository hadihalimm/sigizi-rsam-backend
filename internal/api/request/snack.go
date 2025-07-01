package request

type CreateSnack struct {
	Name string `json:"name" binding:"required" validate:"required"`
}

type UpdateSnack struct {
	Name string `json:"name" binding:"required" validate:"required"`
}

type CreateSnackVariant struct {
	SnackID                    uint   `json:"snackID" binding:"required" validate:"required"`
	Name                       string `json:"name" binding:"required" validate:"required"`
	MealTypeIDs                []uint `json:"mealTypeIDs" binding:"required" validate:"required"`
	DietIDs                    []uint `json:"dietIDs" binding:"required" validate:"required"`
	SnackVariantMaterialUsages []struct {
		FoodMaterialID uint    `json:"foodMaterialID" validate:"required"`
		QuantityUsed   float64 `json:"quantityUsed" validate:"required,gt=0"`
	} `json:"snackVariantMaterialUsages" validate:"required,dive"`
}

type UpdateSnackVariant struct {
	SnackID                    uint   `json:"snackID" binding:"required" validate:"required"`
	Name                       string `json:"name" binding:"required" validate:"required"`
	MealTypeIDs                []uint `json:"mealTypeIDs" binding:"required" validate:"required"`
	DietIDs                    []uint `json:"dietIDs" binding:"required" validate:"required"`
	SnackVariantMaterialUsages []struct {
		FoodMaterialID uint    `json:"foodMaterialID" validate:"required"`
		QuantityUsed   float64 `json:"quantityUsed" validate:"required,gt=0"`
	} `json:"snackVariantMaterialUsages" validate:"required,dive"`
}
