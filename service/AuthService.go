package service

import (
	"strconv"

	"github.com/adenhidayatuloh/Kelompok5_FinalProject3/entity"
	"github.com/adenhidayatuloh/Kelompok5_FinalProject3/pkg/errs"
	taskrepository "github.com/adenhidayatuloh/Kelompok5_FinalProject3/repository/taskRepository"
	"github.com/adenhidayatuloh/Kelompok5_FinalProject3/repository/userrepository"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
	AdminAuthorization() gin.HandlerFunc
	TaskAuthorization() gin.HandlerFunc
}

type authService struct {
	userRepo userrepository.UserRepository
	taskRepo taskrepository.TaskRepository
}

func NewAuthService(userRepo userrepository.UserRepository, taskRepo taskrepository.TaskRepository) AuthService {
	return &authService{userRepo: userRepo, taskRepo: taskRepo}
}

func (a *authService) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearerToken := ctx.GetHeader("Authorization")

		var user entity.User

		if err := user.ValidateToken(bearerToken); err != nil {
			ctx.AbortWithStatusJSON(err.StatusCode(), err)
			return
		}

		result, err := a.userRepo.GetUserByID(user.ID)
		if err != nil {
			ctx.AbortWithStatusJSON(err.StatusCode(), err)
			return
		}

		ctx.Set("userData", result)
		ctx.Next()
	}
}

func (a *authService) AdminAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData, ok := ctx.MustGet("userData").(*entity.User)
		if !ok {
			newError := errs.NewBadRequest("Failed to get user data")
			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
			return
		}

		if userData.Role != "admin" {
			newError := errs.NewUnauthorized("You're not authorized to access this endpoint")
			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
			return
		}

		ctx.Next()
	}
}

func (a *authService) TaskAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData, ok := ctx.MustGet("userData").(*entity.User)
		if !ok {
			newError := errs.NewBadRequest("Failed to get user data")
			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
			return
		}

		taskID := ctx.Param("taskId")
		taskIDUint, err := strconv.ParseUint(taskID, 10, 32)
		if err != nil {
			newError := errs.NewBadRequest("Task id should be an unsigned integer")
			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
			return
		}

		_ = taskIDUint

		task, err2 := a.taskRepo.GetTaskByID(uint(taskIDUint))
		if err2 != nil {
			ctx.AbortWithStatusJSON(err2.StatusCode(), err2)
			return
		}

		if task.UserID != userData.ID {
			newError := errs.NewUnauthorized("You're not authorized to modify this task")
			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
			return
		}

		ctx.Next()
	}
}
