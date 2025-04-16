package request

type CreateRoom struct {
	RoomNumber     string `json:"roomNumber" validate:"required"`
	TreatmentClass string `json:"treatmentClass" validate:"required"`
	RoomTypeID     uint   `json:"roomTypeID" validate:"required"`
}

type UpdateRoom struct {
	RoomNumber     string `json:"roomNumber" validate:"required"`
	TreatmentClass string `json:"treatmentClass" validate:"required"`
	RoomTypeID     uint   `json:"roomTypeID" validate:"required"`
}
