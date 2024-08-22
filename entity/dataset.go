package entity

import (
	"gorm.io/gorm"
)

type Dataset struct {
	gorm.Model
	LocationId     string  `json:"location_id" gorm:"size:32"`
	Name           string  `json:"name" gorm:"size:255"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	BusinessStatus string  `json:"business_status" gorm:"size:255"`
	Kelurahan      string  `json:"kelurahan" gorm:"size:255"`
	Kota           string  `json:"kota" gorm:"size:255"`
	Category       string  `json:"category" gorm:"size:255"`
}
