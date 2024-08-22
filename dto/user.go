package dto

import (
	"mime/multipart"
	"time"
)

type RegisterUserRequest struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Password    string `json:"password" validate:"required,min=8"`
}

type RegisterUserResponse struct {
	Status      int       `json:"status"`
	ID          int       `json:"user_id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginUserResponse struct {
	Status  int    `json:"status"`
	Token   string `json:"token"`
	Role    string `json:"role"`
	User_ID int    `json:"user_id"`
	Name    string `json:"name"`
}

type User struct {
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetUserByIDResponse struct {
	Status      int       `json:"status"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	DokumenID   string    `json:"dokumen_id"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type LogoutResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ResetPasswordRequest struct {
	Email           string `json:"email" validate:"required"`
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required,min=8"`
}

type ResetPasswordResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type CreateDokumenRequest struct {
	FormFile []*multipart.FileHeader
}

type CreateImageResponse struct {
	ID          int       `json:"user_id"`
	DokumenID   string    `json:"dokumen_id"`
	DokumenPath string    `json:"dokumen_path"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetImageByUserIDResponse struct {
	Status      int       `json:"status"`
	ID          int       `json:"user_id"`
	DokumenID   string    `json:"dokumen_id"`
	DokumenPath string    `json:"dokumen_path"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
