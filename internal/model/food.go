package model

type Food struct {
	ID                 uint                `gorm:"primaryKey" json:"id"`
	Name               string              `gorm:"not null" json:"name"`
	FoodMaterialUsages []FoodMaterialUsage `json:"foodMaterialUsages"`
}

type FoodMaterialUsage struct {
	FoodID         uint         `gorm:"primaryKey" json:"foodID"`
	FoodMaterialID uint         `gorm:"primaryKey" json:"foodMaterialID"`
	QuantityUsed   float64      `gorm:"not null" json:"quantityUsed"`
	FoodMaterial   FoodMaterial `json:"foodMaterial"`
}
