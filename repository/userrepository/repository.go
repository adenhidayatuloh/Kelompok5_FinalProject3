package userrepository

import (
	"finalProject3/entity"
	"finalProject3/pkg/errs"
)

type UserRepository interface {
	Register(*entity.User) (*entity.User, errs.MessageErr)
	GetUserByEmail(email string) (*entity.User, errs.MessageErr)
	GetUserByID(id uint) (*entity.User, errs.MessageErr)
	UpdateUser(oldUser *entity.User, newUser *entity.User) (*entity.User, errs.MessageErr)
	DeleteUser(user *entity.User) errs.MessageErr
}
