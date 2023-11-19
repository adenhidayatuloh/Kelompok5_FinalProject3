package service

import (
	"finalProject3/dto"
	"finalProject3/entity"
	"finalProject3/pkg"
	"finalProject3/pkg/errs"
	"finalProject3/repository/userrepository"
	"fmt"
)

type UserService interface {
	Register(payload *dto.RegisterRequest) (*dto.RegisterResponse, errs.MessageErr)
	Login(payload *dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr)
	UpdateUser(user *entity.User, payload *dto.UpdateUserRequest) (*dto.UpdateUserResponse, errs.MessageErr)
	DeleteUser(user *entity.User) (*dto.DeleteUserResponse, errs.MessageErr)
}

type userService struct {
	userRepo userrepository.UserRepository
}

func NewUserService(userRepo userrepository.UserRepository) UserService {
	return &userService{userRepo}
}

func (u *userService) Register(payload *dto.RegisterRequest) (*dto.RegisterResponse, errs.MessageErr) {

	err := pkg.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	user := entity.User{
		FullName: payload.FullName,
		Email:    payload.Email,
		Role:     "member",
		Password: payload.Password,
	}

	if err := user.HashPassword(); err != nil {
		return nil, err
	}

	registeredUser, err := u.userRepo.Register(&user)
	if err != nil {
		return nil, err
	}

	response := &dto.RegisterResponse{
		ID:        registeredUser.ID,
		FullName:  registeredUser.FullName,
		Email:     registeredUser.Email,
		CreatedAt: registeredUser.CreatedAt,
	}

	return response, nil
}

func (u *userService) Login(payload *dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr) {

	err := pkg.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	user, err := u.userRepo.GetUserByEmail(payload.Email)
	if err != nil {
		return nil, err
	}
	fmt.Println(user)

	if err := user.ComparePassword(payload.Password); err != nil {
		return nil, err
	}

	token, err2 := user.CreateToken()
	if err2 != nil {
		return nil, err2
	}

	response := &dto.LoginResponse{Token: token}

	return response, nil
}

func (u *userService) UpdateUser(user *entity.User, payload *dto.UpdateUserRequest) (*dto.UpdateUserResponse, errs.MessageErr) {

	err := pkg.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	newUser := entity.User{
		FullName: payload.FullName,
		Email:    payload.Email,
	}

	updatedUser, err := u.userRepo.UpdateUser(user, &newUser)
	if err != nil {
		return nil, err
	}

	response := &dto.UpdateUserResponse{
		ID:        updatedUser.ID,
		FullName:  updatedUser.FullName,
		Email:     updatedUser.Email,
		UpdatedAt: updatedUser.UpdatedAt,
	}

	return response, nil
}

func (u *userService) DeleteUser(user *entity.User) (*dto.DeleteUserResponse, errs.MessageErr) {
	if err := u.userRepo.DeleteUser(user); err != nil {
		return nil, err
	}

	response := &dto.DeleteUserResponse{
		Message: "Your account has been successfully deleted",
	}

	return response, nil
}
