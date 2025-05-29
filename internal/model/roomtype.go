package model

type RoomType struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"unique;not null" json:"name"`
}

type RSAMRoomType struct {
	Code string `json:"kode_bangsal"`
	Name string `json:"nama_bangsal"`
}
