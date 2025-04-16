package request

type CreateRoom struct {
	RoomNumber     string `json:"roomNumber" binding:"required" validate:"required"`
	TreatmentClass string `json:"treatmentClass" binding:"required" validate:"required"`
	RoomTypeID     uint   `json:"roomTypeID" binding:"required" validate:"required"`
}

type UpdateRoom struct {
	RoomNumber     string `json:"roomNumber" binding:"required" validate:"required"`
	TreatmentClass string `json:"treatmentClass" binding:"required" validate:"required"`
	RoomTypeID     uint   `json:"roomTypeID" binding:"required" validate:"required"`
}
