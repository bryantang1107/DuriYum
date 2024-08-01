package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID     string      `json:"user_id"`
	Price      uint        `json:"price"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID" json:"order_items"`
}
