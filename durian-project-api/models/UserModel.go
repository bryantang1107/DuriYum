package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       string  `gorm:"primaryKey" json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Phone    int     `json:"phone"`
	Unit     string  `json:"unit"`
	Street   string  `json:"street"`
	City     string  `json:"city"`
	State    string  `json:"state"`
	Postcode int     `json:"postcode"`
	Role     string  `json:"role"`
	Orders   []Order `gorm:"foreignKey:UserID" json:"orders"` // One-to-many relationship
}
