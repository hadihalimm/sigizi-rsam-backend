package model

type Allergy struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Code string `gorm:"unique; not null" json:"code"`
	Name string `gorm:"unique; not null" json:"name"`
}
