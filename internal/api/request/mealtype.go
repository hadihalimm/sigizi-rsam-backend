package request

type CreateMealType struct {
	Name string `json:"name" binding:"required" validate:"required"`
}

type UpdateMealType struct {
	Name string `json:"name" binding:"required" validate:"required"`
}
