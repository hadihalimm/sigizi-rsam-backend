package request

type CreateRoom struct {
	RoomNumber     string `json:"room_number" validate:"required"`
	TreatmentClass string `json:"treatment_class" validate:"required"`
	RoomTypeID     uint   `json:"room_type_id" validate:"required"`
}

type UpdateRoom struct {
	RoomNumber     string `json:"room_number" validate:"required"`
	TreatmentClass string `json:"treatment_class" validate:"required"`
	RoomTypeID     uint   `json:"room_type_id" validate:"required"`
}
