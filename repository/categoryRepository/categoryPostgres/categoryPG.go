package categorypostgres

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/adenhidayatuloh/Kelompok5_FinalProject3/entity"
	"github.com/adenhidayatuloh/Kelompok5_FinalProject3/pkg/errs"
	categoryrepository "github.com/adenhidayatuloh/Kelompok5_FinalProject3/repository/categoryRepository"
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

	err := c.db.First(&category, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("Error:", err.Error())
		return nil, errs.NewNotFound(fmt.Sprintf("Category with id %d is not found", id))
	}

	if err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError(fmt.Sprintf("Failed to get category with id %d", id))
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

	err := c.db.Delete(category).Error
	if strings.Contains(err.Error(), gorm.ErrForeignKeyViolated.Error()) {
		return errs.NewForeignkeyViolates(fmt.Sprintf("Category id %d has a reference to the task", category.ID))

	}

	if err != nil {

		log.Println("Error:", err.Error())
		return errs.NewInternalServerError(fmt.Sprintf("Failed to delete category with id %d", category.ID))
	}

	return nil

}
