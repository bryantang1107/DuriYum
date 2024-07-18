package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Unit     string `json:"unit"`
	Street   string `json:"street"`
	City     string `json:"city"`
	State    string `json:"state"`
	PostCode int    `json:"post_code"`
	// 1-M
	Orders []Order `json:"orders"`
}
