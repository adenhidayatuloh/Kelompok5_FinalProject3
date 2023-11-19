package categorypostgres

import (
	"finalProject3/entity"
	"finalProject3/pkg/errs"
	categoryrepository "finalProject3/repository/categoryRepository"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type categoryPG struct {
	db *gorm.DB
}

func NewCategoryPG(db *gorm.DB) categoryrepository.CategoryRepository {
	return &categoryPG{db}
}

func (c *categoryPG) CreateCategory(category *entity.Category) (*entity.Category, errs.MessageErr) {
	if err := c.db.Create(category).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError("Failed to create new category")
	}

	return category, nil
}

func (c *categoryPG) GetAllCategories() ([]entity.Category, errs.MessageErr) {
	var categories []entity.Category

	if err := c.db.Preload("Task").Find(&categories).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError("Failed to get all categories")
	}

	return categories, nil
}

func (c *categoryPG) GetCategoryByID(id uint) (*entity.Category, errs.MessageErr) {
	var category entity.Category

	if err := c.db.First(&category, id).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewNotFound(fmt.Sprintf("Category with %d is not found", id))
	}

	return &category, nil
}

func (c *categoryPG) UpdateCategory(oldCategory *entity.Category, newCategory *entity.Category) (*entity.Category, errs.MessageErr) {
	if err := c.db.Model(oldCategory).Updates(newCategory).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError(fmt.Sprintf("Failed to update category with id %d", oldCategory.ID))
	}

	return oldCategory, nil
}

func (c *categoryPG) DeleteCategory(category *entity.Category) errs.MessageErr {
	if err := c.db.Delete(category).Error; err != nil {
		log.Println("Error:", err.Error())
		return errs.NewInternalServerError(fmt.Sprintf("Failed to delete category with id %d", category.ID))
	}

	return nil
}
