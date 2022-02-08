package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	ID       uint64 `gorm:"primaryKey:autoIncrement" json:"id"`
	Title    string `gorm:"type:varchar(100)" json:"title" binding:"required"`
	Body     string `gorm:"type:varchar(100)" json:"body" binding:"required"`
	AuthorID uint64 `gorm:"not null" json:"-"`
	Author   Author `gorm:"embedded;embeddedPrefix:author_" json:"author"`
}
