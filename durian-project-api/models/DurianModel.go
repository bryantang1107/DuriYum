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
	Name         string `json:"name"`
	Description  string `json:"description"`
	PriceInKG    int    `json:"price_in_kg"`
	PriceInBox   int    `json:"price_in_box"`
	Is_Available bool   `json:"is_available"`
}
