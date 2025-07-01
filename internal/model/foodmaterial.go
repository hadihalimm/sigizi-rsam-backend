package model

type FoodMaterial struct {
	ID              uint    `gorm:"primaryKey" json:"id"`
	Name            string  `gorm:"not null" json:"name"`
	Category        string  `gorm:"not null;index" json:"category"`
	Unit            string  `gorm:"not null" json:"unit"`
	StandardPerMeal float64 `gorm:"not null" json:"standardPerMeal"`
}
