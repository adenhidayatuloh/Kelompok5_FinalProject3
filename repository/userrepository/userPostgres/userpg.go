package userpostgres

import (
	"finalProject3/entity"
	"finalProject3/pkg/errs"
	"finalProject3/repository/userrepository"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type userPG struct {
	db *gorm.DB
}

func NewUserPG(db *gorm.DB) userrepository.UserRepository {
	return &userPG{db}
}

func (u *userPG) Register(user *entity.User) (*entity.User, errs.MessageErr) {

	if err := u.db.Create(user).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError("Failed to register new user")
	}

	return user, nil
}

func (u *userPG) GetUserByEmail(email string) (*entity.User, errs.MessageErr) {
	var user entity.User

	if err := u.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, errs.NewNotFound(fmt.Sprintf("User with email %s is not found", email))
	}

	return &user, nil
}

func (u *userPG) GetUserByID(id uint) (*entity.User, errs.MessageErr) {
	var user entity.User

	if err := u.db.First(&user, id).Error; err != nil {
		return nil, errs.NewNotFound(fmt.Sprintf("User with id %d is not found", id))
	}

	return &user, nil
}

func (u *userPG) UpdateUser(oldUser *entity.User, newUser *entity.User) (*entity.User, errs.MessageErr) {
	if err := u.db.Model(oldUser).Updates(newUser).Error; err != nil {
		return nil, errs.NewInternalServerError(fmt.Sprintf("Failed to update user with id %d", oldUser.ID))
	}

	return oldUser, nil
}

func (u *userPG) DeleteUser(user *entity.User) errs.MessageErr {
	if err := u.db.Delete(user).Error; err != nil {
		log.Println(err.Error())
		return errs.NewInternalServerError(fmt.Sprintf("Failed to delete user with id %d", user.ID))
	}

	return nil
}
