package models

import "gorm.io/gorm"

type PostModel struct {
	gorm.Model
	Name    string `json:"name"`
	Subject string `json:"subject"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
	Message string `json:"message"`
}
