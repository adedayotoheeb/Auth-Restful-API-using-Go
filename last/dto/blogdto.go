package dto

type UserDTO struct {
	Name string `json:"name" form:"name" binding:"required" validate:"min:2"`
	Email string `json:"email" form:"email" binding:"required,email" validate:"email,required"`
	Password string	`json:"password,omitempty" form:"password,omitempty" binding:"required" validate:"min:6"`
}