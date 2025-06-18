package model

import "time"

type Food struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Name          string         `gorm:"not null" json:"name"`
	FoodMaterials []FoodMaterial `gorm:"many2many:food_material_usages;" json:"foodMaterials"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
}

type FoodMaterialUsage struct {
	FoodID         uint    `gorm:"primaryKey" json:"foodID"`
	FoodMaterialID uint    `gorm:"primaryKey" json:"foodMaterialID"`
	QuantityUsed   float64 `gorm:"not null" json:"quantityUsed"`
}
