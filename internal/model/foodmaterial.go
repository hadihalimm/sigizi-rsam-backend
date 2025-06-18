package model

import "time"

type FoodMaterial struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"not null" json:"name"`
	Unit         string    `gorm:"not null" json:"unit"`
	PricePerUnit float64   `gorm:"not null" json:"price_per_unit"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
