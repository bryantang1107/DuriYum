package models

import "gorm.io/gorm"

// Musang King AAA
// Musang King A
// Musang King B
// Black Thorn (D200)
// Traka (竹脚）
// XO
// D24
// Kampung
type Durian struct {
	gorm.Model
	Code         int    `gorm:"primaryKey" json:"code"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	PricePerKg   int    `json:"price_per_kg"`
	PricePerBox  int    `json:"price_per_box"`
	Availability string `json:"availability"`
}
