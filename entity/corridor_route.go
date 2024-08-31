package entity

import "gorm.io/gorm"

type Corridor_Route struct {
	gorm.Model
	Name      string  `json:"name" gorm:"size:255"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Route     string  `json:"route" gorm:"size:50"`
	Direction string  `json:"direction" gorm:"size:50"`
}
