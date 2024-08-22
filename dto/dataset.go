package dto

import (
	"time"
)

// CreateDatasetRequest digunakan untuk membuat dataset baru
type CreateDatasetRequest struct {
	LocationId     string  `json:"location_id" validate:"required"`
	Name           string  `json:"name" validate:"required"`
	Latitude       float64 `json:"latitude" validate:"required"`
	Longitude      float64 `json:"longitude" validate:"required"`
	BusinessStatus string  `json:"business_status" validate:"required"`
	Kelurahan      string  `json:"kelurahan" validate:"required"`
	Kota           string  `json:"kota" validate:"required"`
	Category       string  `json:"category" validate:"required"`
}

// CreateDatasetResponse digunakan untuk merespons pembuatan dataset baru
type CreateDatasetResponse struct {
	Status         int       `json:"status"`
	ID             int       `json:"dataset_id"`
	LocationId     string    `json:"location_id"`
	Name           string    `json:"name"`
	Latitude       float64   `json:"latitude"`
	Longitude      float64   `json:"longitude"`
	BusinessStatus string    `json:"business_status"`
	Kelurahan      string    `json:"kelurahan"`
	Kota           string    `json:"kota"`
	Category       string    `json:"category"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// UpdateDatasetRequest digunakan untuk memperbarui dataset yang ada
type UpdateDatasetRequest struct {
	LocationId     string  `json:"location_id,omitempty"`
	Name           string  `json:"name,omitempty"`
	Latitude       float64 `json:"latitude,omitempty"`
	Longitude      float64 `json:"longitude,omitempty"`
	BusinessStatus string  `json:"business_status,omitempty"`
	Kelurahan      string  `json:"kelurahan,omitempty"`
	Kota           string  `json:"kota,omitempty"`
	Category       string  `json:"category,omitempty"`
}

// UpdateDatasetResponse digunakan untuk merespons pembaruan dataset
type UpdateDatasetResponse struct {
	Status         int       `json:"status"`
	ID             int       `json:"dataset_id"`
	LocationId     string    `json:"location_id"`
	Name           string    `json:"name"`
	Latitude       float64   `json:"latitude"`
	Longitude      float64   `json:"longitude"`
	BusinessStatus string    `json:"business_status"`
	Kelurahan      string    `json:"kelurahan"`
	Kota           string    `json:"kota"`
	Category       string    `json:"category"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// GetDatasetByIDResponse digunakan untuk mendapatkan dataset berdasarkan ID
type GetDatasetByIDResponse struct {
	Status         int       `json:"status"`
	ID             int       `json:"dataset_id"`
	LocationId     string    `json:"location_id"`
	Name           string    `json:"name"`
	Latitude       float64   `json:"latitude"`
	Longitude      float64   `json:"longitude"`
	BusinessStatus string    `json:"business_status"`
	Kelurahan      string    `json:"kelurahan"`
	Kota           string    `json:"kota"`
	Category       string    `json:"category"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// GetDatasetByNameResponse digunakan untuk mendapatkan dataset berdasarkan nama
type GetDatasetByNameResponse struct {
	Status         int       `json:"status"`
	ID             int       `json:"dataset_id"`
	LocationId     string    `json:"location_id"`
	Name           string    `json:"name"`
	Latitude       float64   `json:"latitude"`
	Longitude      float64   `json:"longitude"`
	BusinessStatus string    `json:"business_status"`
	Kelurahan      string    `json:"kelurahan"`
	Kota           string    `json:"kota"`
	Category       string    `json:"category"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type GetDatasetByCategoryResponse struct {
	ID             int       `json:"dataset_id"`
	LocationId     string    `json:"location_id"`
	Name           string    `json:"name"`
	Latitude       float64   `json:"latitude"`
	Longitude      float64   `json:"longitude"`
	BusinessStatus string    `json:"business_status"`
	Kelurahan      string    `json:"kelurahan"`
	Kota           string    `json:"kota"`
	Category       string    `json:"category"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// GetAllDatasetsResponse digunakan untuk merespons informasi semua dataset
type GetAllDatasetsResponse struct {
	Status   int           `json:"status"`
	Message  string        `json:"message"`
	Datasets []DatasetInfo `json:"datasets"`
}

// DatasetInfo digunakan untuk mendeskripsikan setiap dataset dalam GetAllDatasetsResponse
type DatasetInfo struct {
	ID             int       `json:"dataset_id"`
	LocationId     string    `json:"location_id"`
	Name           string    `json:"name"`
	Latitude       float64   `json:"latitude"`
	Longitude      float64   `json:"longitude"`
	BusinessStatus string    `json:"business_status"`
	Kelurahan      string    `json:"kelurahan"`
	Kota           string    `json:"kota"`
	Category       string    `json:"category"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// DeleteDatasetRequest digunakan untuk meminta penghapusan dataset
type DeleteDatasetRequest struct {
	ID int `json:"dataset_id"`
}

// DeleteDatasetResponse digunakan untuk merespons penghapusan dataset
type DeleteDatasetResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
