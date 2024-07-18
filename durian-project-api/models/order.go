package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID     uint        `json:"user_id"`
	OrderItems []OrderItem `json:"order_items"`
	Price      int         `json:"price"`
}
