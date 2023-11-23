package taskpostgres

import (
	"errors"
	"fmt"
	"log"

	"github.com/adenhidayatuloh/Kelompok5_FinalProject3/entity"
	"github.com/adenhidayatuloh/Kelompok5_FinalProject3/pkg/errs"
	taskrepository "github.com/adenhidayatuloh/Kelompok5_FinalProject3/repository/taskRepository"
	"gorm.io/gorm"
)

type taskPG struct {
	db *gorm.DB
}

func NewTaskPG(db *gorm.DB) taskrepository.TaskRepository {
	return &taskPG{db}
}

func (t *taskPG) CreateTask(task *entity.Task) (*entity.Task, errs.MessageErr) {

	if err := t.db.Create(task).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError("Failed to create new Task")
	}

	return task, nil
}

func (t *taskPG) GetAllTasks() ([]entity.Task, errs.MessageErr) {
	var tasks []entity.Task
	if err := t.db.Preload("User").Find(&tasks).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError("Failed to geet all task")
	}
	return tasks, nil
}

func (t *taskPG) GetTaskByID(id uint) (*entity.Task, errs.MessageErr) {
	var task entity.Task

	err := t.db.First(&task, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("Error:", err.Error())
		return nil, errs.NewNotFound(fmt.Sprintf("Task with id %d is not found", id))
	}

	if err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError(fmt.Sprintf("Failed to get task with id %d", id))
	}

	return &task, nil
}

func (t *taskPG) UpdateTask(oldTask *entity.Task, newTask *entity.Task) (*entity.Task, errs.MessageErr) {
	if err := t.db.Model(oldTask).Updates(newTask).Error; err != nil {
		return nil, errs.NewInternalServerError(fmt.Sprintf("Failed to update task with id %d", oldTask.ID))
	}
	return oldTask, nil
}

func (t *taskPG) UpdateTaskStatus(id uint, newStatus bool) (*entity.Task, errs.MessageErr) {
	task, err := t.GetTaskByID(id)
	if err != nil {
		return nil, err
	}

	if err := t.db.Model(&task).Update("Status", newStatus).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError("Failed to update task status")
	}

	return task, nil
}

func (t *taskPG) UpdateTaskCategory(id uint, newCategoryID uint) (*entity.Task, errs.MessageErr) {
	task, err := t.GetTaskByID(id)
	if err != nil {
		return nil, err
	}

	if err := t.db.Model(&task).Update("CategoryID", newCategoryID).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError("Failed to update task category")
	}

	return task, nil
}

func (t *taskPG) DeleteTask(id uint) errs.MessageErr {
	if err := t.db.Delete(&entity.Task{}, id).Error; err != nil {
		log.Println("Error:", err.Error())
		return errs.NewInternalServerError(fmt.Sprintf("Failed to delete Task with id %d", id))
	}

	return nil
}
