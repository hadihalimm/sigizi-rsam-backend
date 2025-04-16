package model

type Room struct {
	ID             uint     `gorm:"primaryKey" json:"id"`
	RoomNumber     string   `gorm:"not null" json:"room_number"`
	TreatmentClass string   `gorm:"not null" json:"treatment_class"`
	RoomTypeID     uint     `gorm:"not null" json:"room_type_id"`
	RoomType       RoomType `json:"room_type"`
}
