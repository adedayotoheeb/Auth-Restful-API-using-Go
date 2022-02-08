package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	ID    uint64 `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"type:varchar(100)" json:"name" binding:"required" validate:"alpha"`
	Email string `gorm:"type:varchar(100),uniqueIndex" json:"email" binding:"required,email" validate:"email,unique"`
}
