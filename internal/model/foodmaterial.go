package model

type FoodMaterial struct {
	ID              uint    `gorm:"primaryKey" json:"id"`
	Name            string  `gorm:"not null" json:"name"`
	Unit            string  `gorm:"not null" json:"unit"`
	StandardPerMeal float64 `gorm:"not null" json:"standard_per_unit"`
}
