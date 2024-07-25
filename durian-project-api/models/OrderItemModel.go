package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID  uint   `json:"order_id"`
	DurianID uint   `json:"durian_id"`
	Quantity int    `json:"quantity"`
	Order    Order  `gorm:"foreignKey:OrderID" json:"order"`
	Durian   Durian `gorm:"foreignKey:DurianID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"durian"`
}
