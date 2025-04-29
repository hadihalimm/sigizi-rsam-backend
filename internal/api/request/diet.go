package request

type CreateDiet struct {
	Code string `json:"code" binding:"required" validate:"required"`
	Name string `json:"name" binding:"required" validate:"required"`
}

type UpdateDiet struct {
	Code string `json:"code" binding:"required" validate:"required"`
	Name string `json:"name" binding:"required" validate:"required"`
}
