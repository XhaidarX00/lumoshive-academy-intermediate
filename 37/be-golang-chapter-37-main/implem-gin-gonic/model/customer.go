package model

type Customer struct {
	ID    int    `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}


