package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string `json:"name"`
	Email       string `json:"email" gorm:"unique"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role" gorm:"default:user"`
	Password    string `json:"password"`
	DokumenID   string `json:"dokumen_id"`
	DokumenPath string `json:"dokumen_path"`
}
