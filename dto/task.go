package dto

import (
	"time"
)

type CreateTaskRequest struct {
	Title       string `json:"title" valid:"required~title cannot be ampty"`
	Description string `json:"description" valid:"required~description cannot be ampty"`
	CategoryID  uint   `json:"category_id" valid:"required~Category id cannot be ampty"`
}

type CreateTaskResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Status      bool      `json:"status"`
	Description string    `json:"description"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type GetAllTasksResponse struct {
	ID          uint         `json:"id"`
	Title       string       `json:"title"`
	Status      bool         `json:"status"`
	Description string       `json:"description"`
	UserID      uint         `json:"user_id"`
	CategoryID  uint         `json:"category_id"`
	CreatedAt   time.Time    `json:"created_at"`
	User        TaskUserData `json:"user"`
}

type UpdateTaskRequest struct {
	Title       string `json:"title" valid:"required~title cannot be ampty"`
	Description string `json:"description" valid:"required~description cannot be ampty"`
}

type UpdateTaskResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateTaskStatusRequest struct {
	Status bool `json:"status" valid:"type(bool)~Status must boolean"`
}

type UpdateTaskCategoryRequest struct {
	CategoryID uint `json:"category_id" valid:"required~category id  cannot be ampty"`
}

type DeleteTaskResponse struct {
	Message string `json:"message"`
}

type TaskUserData struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}
