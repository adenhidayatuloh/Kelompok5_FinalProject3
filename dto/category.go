package dto

import (
	"time"
)

type CreateCategoryRequest struct {
	Type string `json:"type" valid:"required~Type cannot be ampty"`
}

type CreateCategoryResponse struct {
	ID        uint      `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

type GetAllCategoriesResponse struct {
	ID               uint               `json:"id"`
	Type             string             `json:"type"`
	UpdatedAt        time.Time          `json:"updated_at"`
	CreatedAt        time.Time          `json:"created_at"`
	CategoryTaskData []CategoryTaskData `json:"Tasks"`
}

type UpdateCategoryRequest struct {
	Type string `json:"type" valid:"required~Type cannot be ampty"`
}

type UpdateCategoryResponse struct {
	ID        uint      `json:"id"`
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteCategoryResponse struct {
	Message string `json:"message"`
}

type CategoryTaskData struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
