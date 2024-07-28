package models

import "gorm.io/gorm"

// upload durian pic
type Durian struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	PriceInKG   int    `json:"price_in_kg"`
	PriceInBox  int    `json:"price_in_box"`
	IsAvailable bool   `json:"is_available"`
}
