package dto

type AuthorDTO struct {
	Name  string ` json:"name" form:"name" binding:"required" validate:"required,alpha"`
	Email string ` json:"email" form:"email"  binding:"required,email,unique" validate:"required,email,unique"`
}
