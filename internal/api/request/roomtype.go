package request

type CreateRoomType struct {
	Name string `json:"name" binding:"required" validate:"required"`
}

type UpdateRoomType struct {
	Name string `json:"name" binding:"required" validate:"required"`
}
