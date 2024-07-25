package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       string  `gorm:"primaryKey" json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Phone    int     `json:"phone"`
	Unit     string  `json:"unit"`
	Street   string  `json:"street"`
	City     string  `json:"city"`
	State    string  `json:"state"`
	PostCode int     `json:"post_code"`
	Orders   []Order `gorm:"foreignKey:UserID" json:"orders"` // One-to-many relationship
}
