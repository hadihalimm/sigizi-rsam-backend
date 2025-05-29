package model

type RoomType struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Code string `gorm:"unique;not null" json:"code"`
	Name string `gorm:"unique;not null" json:"name"`
}

type SIMRSRoomType struct {
	Code string `json:"kode_bangsal"`
	Name string `json:"nama_bangsal"`
}
