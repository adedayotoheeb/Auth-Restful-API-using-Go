package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductName string `gorm:"type:varchar(100)" json:"product_name"`
	Description string `gorm:"type:varchar(100)" json:"description"`
	Price       uint64 `gorm:"type:varchar(100)" json:"price"`
	Inventory   uint64 `gorm:"type:int" json:"inventory"`
	Status      uint64 `gorm:"type:int" json:"status"`
	UserID      uint64 `gorm:"not null" json:"-"`
	User        User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
