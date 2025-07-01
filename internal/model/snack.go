package model

type Snack struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Name          string         `gorm:"not null" json:"name"` // e.g., "Fried Banana"
	SnackVariants []SnackVariant `json:"snackVariants"`
}

type SnackVariant struct {
	ID                         uint                        `gorm:"primaryKey" json:"id"`
	SnackID                    uint                        `gorm:"not null;index" json:"snackID"`
	Name                       string                      `gorm:"not null" json:"name"`
	MealTypes                  []MealType                  `gorm:"many2many:snack_variant_meal_types;" json:"mealTypes"`
	Diets                      []Diet                      `gorm:"many2many:snack_variant_diets;" json:"diets"`
	SnackVariantMaterialUsages []SnackVariantMaterialUsage `json:"snackVariantMaterialUsages"`
	Snack                      Snack                       `json:"snack"`
	// IsDefault          bool                `gorm:"default:false" json:"isDefault"`
}

type SnackVariantMaterialUsage struct {
	SnackVariantID uint         `gorm:"primaryKey" json:"snackVariantID"`
	FoodMaterialID uint         `gorm:"primaryKey" json:"foodMaterialID"`
	QuantityUsed   float64      `gorm:"not null" json:"quantityUsed"`
	FoodMaterial   FoodMaterial `json:"foodMaterial"`
}
