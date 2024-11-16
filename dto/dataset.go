package dto

import (
	"time"
)

// CreateDatasetRequest digunakan untuk membuat dataset baru
type CreateDatasetRequest struct {
	Name      string  `json:"name" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
	Category  string  `json:"category" validate:"required"`
}

// CreateDatasetResponse digunakan untuk merespons pembuatan dataset baru
type CreateDatasetResponse struct {
	// Status    int       `json:"status"`
	ID        int       `json:"dataset_id"`
	Name      string    `json:"name"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Category  string    `json:"category"`
	Kecamatan string    `json:"kecamatan"`
	Kelurahan string    `json:"kelurahan"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UpdateDatasetRequest digunakan untuk memperbarui dataset yang ada
type UpdateDatasetRequest struct {
	Name      string  `json:"name,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Category  string  `json:"category,omitempty"`
	Kecamatan string  `json:"kecamatan,omitempty"`
	Kelurahan string  `json:"kelurahan,omitempty"`
}

// UpdateDatasetResponse digunakan untuk merespons pembaruan dataset
type UpdateDatasetResponse struct {
	// Status    int       `json:"status"`
	ID        int       `json:"dataset_id"`
	Name      string    `json:"name"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Category  string    `json:"category"`
	Kecamatan string    `json:"kecamatan"`
	Kelurahan string    `json:"kelurahan"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetDatasetByIDResponse digunakan untuk mendapatkan dataset berdasarkan ID
type GetDatasetByIDResponse struct {
	// Status    int       `json:"status"`
	ID        int       `json:"dataset_id"`
	Name      string    `json:"name"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Category  string    `json:"category"`
	Kecamatan string    `json:"kecamatan"`
	Kelurahan string    `json:"kelurahan"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetDatasetByNameResponse digunakan untuk mendapatkan dataset berdasarkan nama
type GetDatasetByNameResponse struct {
	// Status    int       `json:"status"`
	ID        int       `json:"dataset_id"`
	Name      string    `json:"name"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Category  string    `json:"category"`
	Kecamatan string    `json:"kecamatan"`
	Kelurahan string    `json:"kelurahan"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetDatasetByKecamatanResponse digunakan untuk mendapatkan dataset berdasarkan nama
type GetDatasetByKecamatanResponse struct {
	// Status    int       `json:"status"`
	ID        int       `json:"dataset_id"`
	Name      string    `json:"name"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Category  string    `json:"category"`
	Kecamatan string    `json:"kecamatan"`
	Kelurahan string    `json:"kelurahan"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetDatasetByKelurahanResponse digunakan untuk mendapatkan dataset berdasarkan nama
type GetDatasetByKelurahanResponse struct {
	// Status    int       `json:"status"`
	ID        int       `json:"dataset_id"`
	Name      string    `json:"name"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Category  string    `json:"category"`
	Kecamatan string    `json:"kecamatan"`
	Kelurahan string    `json:"kelurahan"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetDatasetByCategoryResponse digunakan untuk mendapatkan dataset berdasarkan nama
type GetDatasetByCategoryResponse struct {
	// Status    int       `json:"status"`
	ID        int       `json:"dataset_id"`
	Name      string    `json:"name"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Category  string    `json:"category"`
	Kecamatan string    `json:"kecamatan"`
	Kelurahan string    `json:"kelurahan"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetAllDatasetsResponse digunakan untuk merespons informasi semua dataset
type GetAllDatasetsResponse struct {
	// Status    int       `json:"status"`
	ID        int       `json:"dataset_id"`
	Name      string    `json:"name"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Category  string    `json:"category"`
	Kecamatan string    `json:"kecamatan"`
	Kelurahan string    `json:"kelurahan"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DatasetInfo digunakan untuk mendeskripsikan setiap dataset dalam GetAllDatasetsResponse
type DatasetInfo struct {
	ID        int       `json:"dataset_id"`
	Name      string    `json:"name"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Category  string    `json:"category"`
	Kecamatan string    `json:"kecamatan"`
	Kelurahan string    `json:"kelurahan"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
