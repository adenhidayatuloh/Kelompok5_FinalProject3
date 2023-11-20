package handler

import (
	"net/http"
	"strconv"

	"github.com/adenhidayatuloh/Kelompok5_FinalProject3/dto"
	"github.com/adenhidayatuloh/Kelompok5_FinalProject3/entity"
	"github.com/adenhidayatuloh/Kelompok5_FinalProject3/pkg/errs"
	"github.com/adenhidayatuloh/Kelompok5_FinalProject3/service"
	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskService service.TaskService
}

func NewTaskHandler(taskService service.TaskService) *TaskHandler {
	return &TaskHandler{taskService}
}

func (t *TaskHandler) CreateTask(ctx *gin.Context) {
	userData, ok := ctx.MustGet("userData").(*entity.User)
	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}
	var requestBody dto.CreateTaskRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	createdTask, err := t.taskService.CreateTask(userData.ID, &requestBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusCreated, createdTask)
}

func (t *TaskHandler) GetAllTasks(ctx *gin.Context) {
	tasks, err := t.taskService.GetAllTasks()
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

func (t *TaskHandler) UpdateTask(ctx *gin.Context) {
	taskID := ctx.Param("taskId")
	taskIDUint, err := strconv.ParseUint(taskID, 10, 32)
	if err != nil {
		errValidation := errs.NewBadRequest("Task id should be in unsigned integer")
		ctx.JSON(errValidation.StatusCode(), errValidation)
		return
	}

	var reqBody dto.UpdateTaskRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		errValidation := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(errValidation.StatusCode(), errValidation)
		return
	}

	updatedTask, errUpdate := t.taskService.UpdateTask(uint(taskIDUint), &reqBody)
	if errUpdate != nil {
		ctx.JSON(errUpdate.StatusCode(), errUpdate)
		return
	}

	ctx.JSON(http.StatusOK, updatedTask)
}

func (t *TaskHandler) UpdateTaskStatus(ctx *gin.Context) {
	taskID := ctx.Param("taskId")
	taskIDUint, err := strconv.ParseUint(taskID, 10, 32)
	if err != nil {
		errValidation := errs.NewBadRequest("Task id should be in unsigned integer")
		ctx.JSON(errValidation.StatusCode(), errValidation)
		return
	}

	var reqBody dto.UpdateTaskStatusRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		validationError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(validationError.StatusCode(), validationError)
		return
	}

	response, errUpdate := t.taskService.UpdateTaskStatus(uint(taskIDUint), &reqBody)
	if errUpdate != nil {
		ctx.JSON(errUpdate.StatusCode(), errUpdate)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (t *TaskHandler) UpdateTaskCategory(ctx *gin.Context) {
	taskID := ctx.Param("taskId")
	taskIDUint, err := strconv.ParseUint(taskID, 10, 32)
	if err != nil {
		errValidation := errs.NewBadRequest("Task id should be in unsigned integer")
		ctx.JSON(errValidation.StatusCode(), errValidation)
		return
	}

	var reqBody dto.UpdateTaskCategoryRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		validationError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(validationError.StatusCode(), validationError)
		return
	}

	updatedCategory, errUpdate := t.taskService.UpdateTaskCategory(uint(taskIDUint), &reqBody)
	if errUpdate != nil {
		ctx.JSON(errUpdate.StatusCode(), errUpdate)
		return
	}

	ctx.JSON(http.StatusOK, updatedCategory)
}

func (t *TaskHandler) DeleteTask(ctx *gin.Context) {
	taskID := ctx.Param("taskId")
	taskIDUint, err := strconv.ParseUint(taskID, 10, 32)
	if err != nil {
		newError := errs.NewBadRequest("Task id should be in unsigned integer")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	response, err2 := t.taskService.DeleteTask(uint(taskIDUint))
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
