package domain

type RegisterUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required"`
}

type UpdateUser struct {
	id   int
	Name string
}
