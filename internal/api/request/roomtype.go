package request

type CreateRoomType struct {
	Name string `json:"name" validate:"required"`
}

type UpdateRoomType struct {
	Name string `json:"name" validate:"required"`
}
