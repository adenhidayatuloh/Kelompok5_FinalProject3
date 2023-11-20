package handler

import (
	"net/http"

	"github.com/adenhidayatuloh/glng_ks08_Kelompok5_final_Project_3/dto"
	"github.com/adenhidayatuloh/glng_ks08_Kelompok5_final_Project_3/entity"
	"github.com/adenhidayatuloh/glng_ks08_Kelompok5_final_Project_3/pkg/errs"
	"github.com/adenhidayatuloh/glng_ks08_Kelompok5_final_Project_3/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService}
}

// ==================================================

func (u *UserHandler) Register(ctx *gin.Context) {
	var requestBody dto.RegisterRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	registeredUser, err := u.userService.Register(&requestBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusCreated, registeredUser)
}

func (u *UserHandler) Login(ctx *gin.Context) {
	var requestBody dto.LoginRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	token, err := u.userService.Login(&requestBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, token)
}

func (u *UserHandler) UpdateUser(ctx *gin.Context) {
	var requestBody dto.UpdateUserRequest
	userData, ok := ctx.MustGet("userData").(*entity.User)

	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	updatedUser, err := u.userService.UpdateUser(userData, &requestBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, updatedUser)
}

func (u *UserHandler) DeleteUser(ctx *gin.Context) {
	userData, ok := ctx.MustGet("userData").(*entity.User)
	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	response, err := u.userService.DeleteUser(userData)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
