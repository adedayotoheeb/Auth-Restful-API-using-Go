package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint64 `gorm:"primaryKey:autoI             ncrement" json:"id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Email    string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string `gorm:"->;<-;not null" json:"-"`
	Token    string `gorm:"-" json:"token,omitempty"`
}
