package model

import "time"

type Room struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Code           string    `gorm:"unique;not null" json:"code"`
	Name           string    `gorm:"not null" json:"name"`
	TreatmentClass string    `gorm:"not null" json:"treatmentClass"`
	RoomTypeID     uint      `gorm:"not null;index" json:"roomTypeID"`
	RoomType       RoomType  `json:"roomType"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type SIMRSRoom struct {
	Code           string `json:"kode"`
	Name           string `json:"nama_kamar"`
	ClassID        string `json:"kelas_id"`
	TreatmentClass string `json:"kelas_layanan"`
}
