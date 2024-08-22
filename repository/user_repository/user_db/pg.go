package user_db

import (
	"15min-city/entity"
	"15min-city/pkg/errs"
	"15min-city/repository/user_repository"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user_repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Register(user entity.User) (*entity.User, errs.ErrMessage) {
	result := u.db.Create(&user)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errs.NewBadRequestError("Email already exists")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}

	return &user, nil
}

func (u *userRepository) GetUserByID(id int) (*entity.User, errs.ErrMessage) {
	var user entity.User

	result := u.db.First(&user, id)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewBadRequestError("User not found")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}

	return &user, nil
}

func (u *userRepository) GetUserByEmail(email string) (*entity.User, errs.ErrMessage) {
	var user entity.User

	result := u.db.Where("email = ?", email).First(&user)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewBadRequestError("Email not found")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}

	return &user, nil
}

func (u *userRepository) Update(userUpdate entity.User) errs.ErrMessage {
	result := u.db.Model(&entity.User{}).Where("id = ?", userUpdate.ID).Update("password", userUpdate.Password)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewBadRequestError("User not found")
		}

		return errs.NewInternalServerError("Something went wrong")
	}

	return nil
}

func (u *userRepository) CreateImage(ctx context.Context, id int, img entity.User) (*entity.User, errs.ErrMessage) {
	result := u.db.Model(&entity.User{}).Where("id = ?", id).Updates(img)

	if err := result.Error; err != nil {
		return nil, errs.NewInternalServerError("Something went wrong")
	}

	var updatedImg entity.User
	if err := u.db.Where("id = ?", id).First(&updatedImg).Error; err != nil {
		return nil, errs.NewInternalServerError("Error retrieving updated image user")
	}

	return &updatedImg, nil
}

func (u *userRepository) GetImageByUser(ctx context.Context, id int, imgID int) (*entity.User, errs.ErrMessage) {
	var dokumen entity.User

	result := u.db.Where("id = ? AND dokumen_id= ?", id, imgID).First(&dokumen)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewBadRequestError("Dokumen not found")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}

	return &dokumen, nil
}

func (u *userRepository) GetUserWhatsapp(id int) (string, errs.ErrMessage) {
	var user entity.User

	result := u.db.Select("phone_number").Where("id = ?", id).First(&user)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errs.NewBadRequestError("Whatsapp not found")
		}

		return "", errs.NewInternalServerError("Something went wrong")
	}

	return user.PhoneNumber, nil
}

func (u *userRepository) GetUserEmail(id int) (string, errs.ErrMessage) {
	var user entity.User

	result := u.db.Select("email").Where("id = ?", id).First(&user)
	fmt.Println("Ini dari repo", result)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errs.NewBadRequestError("Whatsapp not found")
		}

		return "", errs.NewInternalServerError("Something went wrong")
	}

	return user.Email, nil
}
