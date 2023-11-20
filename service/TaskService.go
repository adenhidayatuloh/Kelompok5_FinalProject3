package service

import (
	"finalProject3/dto"
	"finalProject3/entity"
	"finalProject3/pkg"
	"finalProject3/pkg/errs"
	categoryrepository "finalProject3/repository/categoryRepository"
	taskrepository "finalProject3/repository/taskRepository"
)

type TaskService interface {
	CreateTask(userID uint, payload *dto.CreateTaskRequest) (*dto.CreateTaskResponse, errs.MessageErr)
	GetAllTasks() ([]dto.GetAllTasksResponse, errs.MessageErr)
	UpdateTask(id uint, payload *dto.UpdateTaskRequest) (*dto.UpdateTaskResponse, errs.MessageErr)
	UpdateTaskStatus(id uint, payload *dto.UpdateTaskStatusRequest) (*dto.UpdateTaskResponse, errs.MessageErr)
	UpdateTaskCategory(id uint, payload *dto.UpdateTaskCategoryRequest) (*dto.UpdateTaskResponse, errs.MessageErr)
	DeleteTask(id uint) (*dto.DeleteTaskResponse, errs.MessageErr)
}

type taskService struct {
	taskRepo     taskrepository.TaskRepository
	categoryRepo categoryrepository.CategoryRepository
}

func NewTaskService(taskRepo taskrepository.TaskRepository, categoryRepo categoryrepository.CategoryRepository) TaskService {
	return &taskService{taskRepo, categoryRepo}
}

func (t *taskService) CreateTask(UserID uint, payload *dto.CreateTaskRequest) (*dto.CreateTaskResponse, errs.MessageErr) {

	err := pkg.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	task := entity.Task{
		Title:       payload.Title,
		Description: payload.Description,
		CategoryID:  payload.CategoryID,
		UserID:      UserID,
		Status:      false,
	}

	if _, checkCategoryErr := t.categoryRepo.GetCategoryByID(task.CategoryID); checkCategoryErr != nil {
		return nil, checkCategoryErr
	}

	createdTask, err := t.taskRepo.CreateTask(&task)
	if err != nil {
		return nil, err
	}

	response := &dto.CreateTaskResponse{
		ID:          createdTask.ID,
		Title:       createdTask.Title,
		Status:      createdTask.Status,
		Description: createdTask.Description,
		UserID:      createdTask.UserID,
		CategoryID:  createdTask.CategoryID,
		CreatedAt:   createdTask.CreatedAt,
	}

	return response, nil
}

func (t *taskService) GetAllTasks() ([]dto.GetAllTasksResponse, errs.MessageErr) {
	tasks, err := t.taskRepo.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := []dto.GetAllTasksResponse{}

	for _, task := range tasks {
		response = append(response, dto.GetAllTasksResponse{
			ID:          task.ID,
			Title:       task.Title,
			Status:      task.Status,
			Description: task.Description,
			UserID:      task.UserID,
			CategoryID:  task.CategoryID,
			CreatedAt:   task.CreatedAt,
			User: dto.TaskUserData{
				ID:       task.User.ID,
				Email:    task.User.Email,
				FullName: task.User.FullName,
			},
		})

	}
	return response, nil
}

func (t *taskService) UpdateTask(id uint, payload *dto.UpdateTaskRequest) (*dto.UpdateTaskResponse, errs.MessageErr) {

	err := pkg.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	oldTask, err := t.taskRepo.GetTaskByID(id)
	if err != nil {
		return nil, err
	}

	newTask := entity.Task{
		Title:       payload.Title,
		Description: payload.Description,
	}

	updatedTask, err2 := t.taskRepo.UpdateTask(oldTask, &newTask)
	if err2 != nil {
		return nil, err2
	}

	response := &dto.UpdateTaskResponse{
		ID:          updatedTask.ID,
		Title:       updatedTask.Title,
		Description: updatedTask.Description,
		Status:      updatedTask.Status,
		UserID:      updatedTask.UserID,
		CategoryID:  updatedTask.CategoryID,
		UpdatedAt:   updatedTask.UpdatedAt,
	}

	return response, nil
}

func (t *taskService) UpdateTaskStatus(id uint, payload *dto.UpdateTaskStatusRequest) (*dto.UpdateTaskResponse, errs.MessageErr) {

	err := pkg.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	result, err := t.taskRepo.UpdateTaskStatus(id, payload.Status)
	if err != nil {
		return nil, err
	}

	response := &dto.UpdateTaskResponse{
		ID:          result.ID,
		Title:       result.Title,
		Description: result.Description,
		Status:      result.Status,
		UserID:      result.UserID,
		CategoryID:  result.CategoryID,
		UpdatedAt:   result.UpdatedAt,
	}

	return response, nil
}

func (t *taskService) UpdateTaskCategory(id uint, payload *dto.UpdateTaskCategoryRequest) (*dto.UpdateTaskResponse, errs.MessageErr) {

	err := pkg.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	if _, checkCategory := t.categoryRepo.GetCategoryByID(payload.CategoryID); checkCategory != nil {
		return nil, checkCategory
	}

	updatedCategory, err := t.taskRepo.UpdateTaskCategory(id, payload.CategoryID)
	if err != nil {
		return nil, err
	}

	response := &dto.UpdateTaskResponse{
		ID:          updatedCategory.ID,
		Title:       updatedCategory.Title,
		Description: updatedCategory.Description,
		Status:      updatedCategory.Status,
		UserID:      updatedCategory.UserID,
		CategoryID:  updatedCategory.CategoryID,
		UpdatedAt:   updatedCategory.UpdatedAt,
	}

	return response, nil
}

func (t *taskService) DeleteTask(id uint) (*dto.DeleteTaskResponse, errs.MessageErr) {
	if err := t.taskRepo.DeleteTask(id); err != nil {
		return nil, err
	}

	response := &dto.DeleteTaskResponse{
		Message: "Task has been successfully deleted",
	}

	return response, nil
}
