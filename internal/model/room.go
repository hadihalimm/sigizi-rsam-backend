package model

import "time"

type Room struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	RoomNumber     string    `gorm:"not null" json:"roomNumber"`
	TreatmentClass string    `gorm:"not null" json:"treatmentClass"`
	RoomTypeID     uint      `gorm:"not null;index" json:"roomTypeID"`
	RoomType       RoomType  `json:"roomType"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
