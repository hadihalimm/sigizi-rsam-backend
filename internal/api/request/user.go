package request

type UpdateUser struct {
	Username string `json:"username" validate:"required,min=4,max=32"`
	Name     string `json:"name" validate:"required,min=1,max=64"`
	Role     string `json:"role" validate:"required,min=1,max=32"`
}

type UpdatePassword struct {
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type UpdateName struct {
	Name string `json:"name" validate:"required,min=1,max=32"`
}
