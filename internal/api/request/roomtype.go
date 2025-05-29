package request

type CreateRoomType struct {
	Name string `json:"name" binding:"required" validate:"required"`
	Code string `json:"code" binding:"required" validate:"required"`
}

type UpdateRoomType struct {
	Name string `json:"name" binding:"required" validate:"required"`
	Code string `json:"code" binding:"required" validate:"required"`
}
