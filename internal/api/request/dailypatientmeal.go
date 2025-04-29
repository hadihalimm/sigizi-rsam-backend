package request

type CreateDailyPatientMeal struct {
	PatientID  uint   `json:"patientID" binding:"required" validate:"required"`
	RoomID     uint   `json:"roomID" binding:"required" validate:"required"`
	MealTypeID uint   `json:"mealTypeID" binding:"required" validate:"required"`
	Notes      string `json:"notes"`
	DietIDs    []uint `json:"dietIDs"`
}

type UpdateDailyPatientMeal struct {
	PatientID  uint   `json:"patientID" binding:"required" validate:"required"`
	RoomID     uint   `json:"roomID" binding:"required" validate:"required"`
	MealTypeID uint   `json:"mealTypeID" binding:"required" validate:"required"`
	Notes      string `json:"notes"`
	DietIDs    []uint `json:"dietIDs"`
}
