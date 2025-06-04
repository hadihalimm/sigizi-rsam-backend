package request

type CreateRoom struct {
	Name           string `json:"name" binding:"required" validate:"required"`
	Code           string `json:"code" binding:"required" validate:"required"`
	ClassID        string `json:"treatment_class_id" binding:"required" validate:"required"`
	TreatmentClass string `json:"treatmentClass" binding:"required" validate:"required"`
	RoomTypeID     uint   `json:"roomTypeID" binding:"required" validate:"required"`
}

type UpdateRoom struct {
	Name           string `json:"name" binding:"required" validate:"required"`
	Code           string `json:"code" binding:"required" validate:"required"`
	ClassID        string `json:"treatment_class_id" binding:"required" validate:"required"`
	TreatmentClass string `json:"treatmentClass" binding:"required" validate:"required"`
	RoomTypeID     uint   `json:"roomTypeID" binding:"required" validate:"required"`
}
