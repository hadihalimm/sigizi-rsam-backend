package model

type Diet struct {
	ID   uint   `gorm:"primaryKey"`
	Code string `gorm:"unique; not null" json:"code"`
	Name string `gorm:"unique; not null" json:"name"`
}
