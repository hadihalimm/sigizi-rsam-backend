package request

type Register struct {
	Username string `json:"username" validate:"required,min=4,max=32"`
	Name     string `json:"name" validate:"required,min=1,max=64"`
	Password string `json:"password" validate:"required,min=8,max=32"`
	Role     string `json:"role" validate:"required,min=1,max=32"`
}

type SignIn struct {
	Username string `json:"username" validate:"required,min=4,max=32"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}
