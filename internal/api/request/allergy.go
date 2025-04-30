package request

type CreateAllergy struct {
	Code string `json:"code" binding:"required" validate:"required"`
	Name string `json:"name" binding:"required" validate:"required"`
}

type UpdateAllergy struct {
	Code string `json:"code" binding:"required" validate:"required"`
	Name string `json:"name" binding:"required" validate:"required"`
}
