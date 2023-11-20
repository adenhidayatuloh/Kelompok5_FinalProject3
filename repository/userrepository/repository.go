package userrepository

import (
	"github.com/adenhidayatuloh/glng_ks08_Kelompok5_final_Project_3/entity"
	"github.com/adenhidayatuloh/glng_ks08_Kelompok5_final_Project_3/pkg/errs"
)

type UserRepository interface {
	Register(*entity.User) (*entity.User, errs.MessageErr)
	GetUserByEmail(email string) (*entity.User, errs.MessageErr)
	GetUserByID(id uint) (*entity.User, errs.MessageErr)
	UpdateUser(oldUser *entity.User, newUser *entity.User) (*entity.User, errs.MessageErr)
	DeleteUser(user *entity.User) errs.MessageErr
}
