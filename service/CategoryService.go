package service

import (
	"github.com/adenhidayatuloh/Kelompok5_FinalProject3/dto"
	"github.com/adenhidayatuloh/Kelompok5_FinalProject3/entity"
	"github.com/adenhidayatuloh/Kelompok5_FinalProject3/pkg"
	"github.com/adenhidayatuloh/Kelompok5_FinalProject3/pkg/errs"
	categoryrepository "github.com/adenhidayatuloh/Kelompok5_FinalProject3/repository/categoryRepository"
)

type CategoryService interface {
	CreateCategory(payload *dto.CreateCategoryRequest) (
		*dto.CreateCategoryResponse,
		errs.MessageErr,
	)

	GetAllCategories() ([]dto.GetAllCategoriesResponse, errs.MessageErr)

	UpdateCategory(id uint, payload *dto.UpdateCategoryRequest) (
		*dto.UpdateCategoryResponse,
		errs.MessageErr,
	)

	DeleteCategory(id uint) (*dto.DeleteCategoryResponse, errs.MessageErr)
}

type categoryService struct {
	categoryRepo categoryrepository.CategoryRepository
}

func NewCategoryService(
	categoryRepo categoryrepository.CategoryRepository,

) CategoryService {
	return &categoryService{categoryRepo}
}

func (c *categoryService) CreateCategory(payload *dto.CreateCategoryRequest) (*dto.CreateCategoryResponse, errs.MessageErr) {

	err := pkg.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	category := entity.Category{
		Type: payload.Type,
	}

	createdCategory, err := c.categoryRepo.CreateCategory(&category)
	if err != nil {
		return nil, err
	}

	response := &dto.CreateCategoryResponse{
		ID:        createdCategory.ID,
		Type:      createdCategory.Type,
		CreatedAt: createdCategory.CreatedAt,
	}

	return response, nil
}

func (c *categoryService) GetAllCategories() ([]dto.GetAllCategoriesResponse, errs.MessageErr) {

	AllCategories, err := c.categoryRepo.GetAllCategories()

	if err != nil {
		return nil, err
	}

	var AllCategoriesTask []dto.GetAllCategoriesResponse

	for _, eachCategories := range AllCategories {

		TaskResult := []dto.CategoryTaskData{}

		CategoryResult := dto.GetAllCategoriesResponse{

			ID:        eachCategories.ID,
			Type:      eachCategories.Type,
			UpdatedAt: eachCategories.UpdatedAt,
			CreatedAt: eachCategories.CreatedAt,
		}

		for _, eachTask := range eachCategories.Task {

			Task := dto.CategoryTaskData{
				ID:          eachTask.ID,
				Title:       eachTask.Title,
				Description: eachTask.Description,
				UserID:      eachTask.UserID,
				CategoryID:  eachTask.CategoryID,
				CreatedAt:   eachTask.CreatedAt,
				UpdatedAt:   eachTask.UpdatedAt,
			}
			TaskResult = append(TaskResult, Task)

		}

		CategoryResult.CategoryTaskData = TaskResult

		AllCategoriesTask = append(AllCategoriesTask, CategoryResult)

	}

	return AllCategoriesTask, nil

}

func (c *categoryService) UpdateCategory(id uint, payload *dto.UpdateCategoryRequest) (*dto.UpdateCategoryResponse, errs.MessageErr) {

	err := pkg.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	oldCategory, err := c.categoryRepo.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}

	newCategory := entity.Category{
		Type: payload.Type,
	}

	updatedCategory, err2 := c.categoryRepo.UpdateCategory(oldCategory, &newCategory)
	if err2 != nil {
		return nil, err2
	}

	response := &dto.UpdateCategoryResponse{
		ID:        updatedCategory.ID,
		Type:      updatedCategory.Type,
		UpdatedAt: updatedCategory.UpdatedAt,
	}

	return response, nil
}

func (c *categoryService) DeleteCategory(id uint) (*dto.DeleteCategoryResponse, errs.MessageErr) {
	category, err := c.categoryRepo.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}

	if err := c.categoryRepo.DeleteCategory(category); err != nil {
		return nil, err
	}

	response := &dto.DeleteCategoryResponse{
		Message: "Category has been successfully deleted",
	}

	return response, nil
}
