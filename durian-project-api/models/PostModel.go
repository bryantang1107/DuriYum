package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Subject string `json:"subject"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
	Message string `json:"message"`
	UserID  string `json:"user_id"`
	User    User   `gorm:"foreignKey:UserID" json:"user"`
}
