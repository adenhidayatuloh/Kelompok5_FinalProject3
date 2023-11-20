package taskrepository

import (
	"github.com/adenhidayatuloh/glng_ks08_Kelompok5_final_Project_3/entity"
	"github.com/adenhidayatuloh/glng_ks08_Kelompok5_final_Project_3/pkg/errs"
)

type TaskRepository interface {
	CreateTask(task *entity.Task) (*entity.Task, errs.MessageErr)
	GetAllTasks() ([]entity.Task, errs.MessageErr)
	GetTaskByID(id uint) (*entity.Task, errs.MessageErr)
	UpdateTask(oldTask *entity.Task, newTask *entity.Task) (*entity.Task, errs.MessageErr)
	UpdateTaskStatus(id uint, newStatus bool) (*entity.Task, errs.MessageErr)
	UpdateTaskCategory(id uint, newCategoryID uint) (*entity.Task, errs.MessageErr)
	DeleteTask(id uint) errs.MessageErr
}
