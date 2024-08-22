package service

import (
	"15min-city/dto"
	"15min-city/entity"
	"15min-city/pkg/errs"
	"15min-city/pkg/helpers"
	"15min-city/repository/user_repository"
	"context"
	"io"
	"net/http"
	"os"
	"strings"
)

type userService struct {
	userRepo user_repository.UserRepository
}

type UserService interface {
	Register(userPayload dto.RegisterUserRequest) (*dto.RegisterUserResponse, errs.ErrMessage)
	Login(userCredentials dto.LoginUserRequest) (*dto.LoginUserResponse, errs.ErrMessage)
	GetUserByID(id int) (*dto.GetUserByIDResponse, errs.ErrMessage)
	ResetPassword(userPayload dto.ResetPasswordRequest) (*dto.ResetPasswordResponse, errs.ErrMessage)
	CreateImage(ctx context.Context, id int, dokumenPayload dto.CreateDokumenRequest) ([]dto.CreateImageResponse, errs.ErrMessage)
	GetImageByUser(ctx context.Context, id int, imgID int) *dto.GetImageByUserIDResponse
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) Register(userPayload dto.RegisterUserRequest) (*dto.RegisterUserResponse, errs.ErrMessage) {
	user := entity.User{
		Name:        userPayload.Name,
		Email:       userPayload.Email,
		PhoneNumber: userPayload.PhoneNumber,
	}

	hashedPassword, err := helpers.HashPassword(userPayload.Password)

	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword

	createdUser, err := u.userRepo.Register(user)

	if err != nil {
		return nil, err
	}

	response := dto.RegisterUserResponse{
		Status:      http.StatusCreated,
		ID:          int(createdUser.ID),
		Name:        createdUser.Name,
		Email:       createdUser.Email,
		PhoneNumber: createdUser.PhoneNumber,
		Role:        createdUser.Role,
		CreatedAt:   createdUser.CreatedAt,
		UpdatedAt:   createdUser.UpdatedAt,
	}

	return &response, nil
}

func (u *userService) GetUserByID(id int) (*dto.GetUserByIDResponse, errs.ErrMessage) {
	user, err := u.userRepo.GetUserByID(id)

	if err != nil {
		return nil, err
	}

	response := dto.GetUserByIDResponse{
		Status:      http.StatusOK,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Role:        user.Role,
		DokumenID:   user.DokumenID,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}

	return &response, nil
}

func (u *userService) Login(userCredentials dto.LoginUserRequest) (*dto.LoginUserResponse, errs.ErrMessage) {
	user, err := u.userRepo.GetUserByEmail(userCredentials.Email)

	if err != nil {
		return nil, err
	}

	if err := helpers.ComparePassword(user.Password, userCredentials.Password); err != nil {
		return nil, err
	}

	token, err := helpers.GenerateToken(int(user.ID), user.Name, user.Email, user.Role)

	if err != nil {
		return nil, err
	}

	response := dto.LoginUserResponse{
		Status:  http.StatusOK,
		Token:   token,
		Role:    user.Role,
		User_ID: int(user.ID),
		Name:    user.Name,
	}

	return &response, nil
}

func (u *userService) ResetPassword(userPayload dto.ResetPasswordRequest) (*dto.ResetPasswordResponse, errs.ErrMessage) {
	user, err := u.userRepo.GetUserByEmail(userPayload.Email)

	if err != nil {
		return nil, err
	}

	if err := helpers.ComparePassword(user.Password, userPayload.CurrentPassword); err != nil {
		return nil, err
	}

	newPassword, err := helpers.HashPassword(userPayload.NewPassword)

	if err != nil {
		return nil, err
	}

	user.Email = userPayload.Email
	user.Password = newPassword

	err = u.userRepo.Update(*user)

	if err != nil {
		return nil, err
	}

	response := dto.ResetPasswordResponse{
		Status:  http.StatusOK,
		Message: "Password has been changed successfully",
	}

	return &response, nil
}

func (u *userService) CreateImage(ctx context.Context, id int, dokumenPayload dto.CreateDokumenRequest) ([]dto.CreateImageResponse, errs.ErrMessage) {
	var dokumenResponses []dto.CreateImageResponse

	for _, dokumen := range dokumenPayload.FormFile {
		file, _ := dokumen.Open()

		tempFile, err := os.CreateTemp("public/imageprofile", "image-*.png")
		if err != nil {
			panic(err)
		}
		defer tempFile.Close()

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		tempFile.Write(fileBytes)

		fileName := tempFile.Name()
		newFileName := strings.Split(fileName, "\\")

		dokumen := entity.User{
			DokumenID:   helpers.GenerateId(),
			DokumenPath: newFileName[1],
		}

		createdDokumen, err := u.userRepo.CreateImage(ctx, id, dokumen)

		response := dto.CreateImageResponse{
			ID:          int(createdDokumen.ID),
			DokumenID:   createdDokumen.DokumenID,
			DokumenPath: createdDokumen.DokumenPath,
			CreatedAt:   createdDokumen.CreatedAt,
			UpdatedAt:   createdDokumen.UpdatedAt,
		}

		dokumenResponses = append(dokumenResponses, response)
	}

	return dokumenResponses, nil
}

func (u *userService) GetImageByUser(ctx context.Context, id int, imgID int) *dto.GetImageByUserIDResponse {
	dokumen, err := u.userRepo.GetImageByUser(ctx, id, imgID)
	if err != nil {
		return nil
	}

	return &dto.GetImageByUserIDResponse{
		ID:          int(dokumen.ID),
		DokumenID:   dokumen.DokumenID,
		DokumenPath: dokumen.DokumenPath,
		CreatedAt:   dokumen.CreatedAt,
		UpdatedAt:   dokumen.UpdatedAt,
	}
}
