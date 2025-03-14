package entity

import (
	"gorm.io/gorm"
)

type Dataset struct {
	gorm.Model
	Name      string  `json:"name" gorm:"size:255"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Category  string  `json:"category" gorm:"size:255"`
	Kecamatan string  `json:"kecamatan" gorm:"size:255"`
	Kelurahan string  `json:"kelurahan" gorm:"size:255"`
	Distance  float64 `json:"distance"`
}
