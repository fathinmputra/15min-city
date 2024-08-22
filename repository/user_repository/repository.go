package user_repository

import (
	"15min-city/entity"
	"15min-city/pkg/errs"
	"context"
)

type UserRepository interface {
	Register(user entity.User) (*entity.User, errs.ErrMessage)
	GetUserByID(id int) (*entity.User, errs.ErrMessage)
	GetUserByEmail(email string) (*entity.User, errs.ErrMessage)
	Update(userUpdate entity.User) errs.ErrMessage
	CreateImage(ctx context.Context, imgID int, img entity.User) (*entity.User, errs.ErrMessage)
	GetImageByUser(ctx context.Context, id int, imgID int) (*entity.User, errs.ErrMessage)
	GetUserWhatsapp(id int) (string, errs.ErrMessage)
	GetUserEmail(id int) (string, errs.ErrMessage)
}
