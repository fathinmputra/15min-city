package dto

import (
	"time"
)

// CreateCorridorRouteRequest digunakan untuk membuat rute koridor baru
type CreateCorridorRouteRequest struct {
	Name      string  `json:"name" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
	Route     string  `json:"route" validate:"required"`
	Direction string  `json:"direction" validate:"required"`
}

// CreateCorridorRouteResponse digunakan untuk merespons pembuatan rute koridor baru
type CreateCorridorRouteResponse struct {
	Status    int       `json:"status"`
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Route     string    `json:"route"`
	Direction string    `json:"direction"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UpdateCorridorRouteRequest digunakan untuk memperbarui rute koridor yang ada
type UpdateCorridorRouteRequest struct {
	Name      string  `json:"name,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Route     string  `json:"route,omitempty"`
	Direction string  `json:"direction,omitempty"`
}

// UpdateCorridorRouteResponse digunakan untuk merespons pembaruan rute koridor
type UpdateCorridorRouteResponse struct {
	Status    int       `json:"status"`
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Route     string    `json:"route"`
	Direction string    `json:"direction"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetCorridorRouteByIDResponse digunakan untuk mendapatkan rute koridor berdasarkan ID
type GetCorridorRouteByIDResponse struct {
	Status    int       `json:"status"`
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Route     string    `json:"route"`
	Direction string    `json:"direction"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetCorridorRouteByNameResponse digunakan untuk mendapatkan rute koridor berdasarkan nama
type GetCorridorRouteByNameResponse struct {
	Status    int       `json:"status"`
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Route     string    `json:"route"`
	Direction string    `json:"direction"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetCorridorRouteByRouteResponse digunakan untuk mendapatkan rute koridor berdasarkan nama
type GetCorridorRouteByRouteResponse struct {
	Status    int       `json:"status"`
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Route     string    `json:"route"`
	Direction string    `json:"direction"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetCorridorRouteByDirectionResponse digunakan untuk mendapatkan rute koridor berdasarkan nama
type GetCorridorRouteByDirectionResponse struct {
	Status    int       `json:"status"`
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Route     string    `json:"route"`
	Direction string    `json:"direction"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetAllCorridorRoutesResponse digunakan untuk merespons informasi semua rute koridor
type GetAllCorridorRoutesResponse struct {
	Status         int                 `json:"status"`
	Message        string              `json:"message"`
	CorridorRoutes []CorridorRouteInfo `json:"rute_koridors"`
}

// CorridorRouteInfo digunakan untuk mendeskripsikan setiap rute koridor dalam GetAllCorridorRoutesResponse
type CorridorRouteInfo struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Route     string    `json:"route"`
	Direction string    `json:"direction"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DeleteCorridorRouteRequest digunakan untuk meminta penghapusan rute koridor
type DeleteCorridorRouteRequest struct {
	ID uint `json:"id"`
}

// DeleteCorridorRouteResponse digunakan untuk merespons penghapusan rute koridor
type DeleteCorridorRouteResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
