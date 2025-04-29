package request

type CreateMealType struct {
	Code string `json:"code" binding:"required" validate:"required"`
	Name string `json:"name" binding:"required" validate:"required"`
}

type UpdateMealType struct {
	Code string `json:"code" binding:"required" validate:"required"`
	Name string `json:"name" binding:"required" validate:"required"`
}
