package categoryrepository

import (
	"github.com/adenhidayatuloh/Kelompok5_FinalProject3/entity"
	"github.com/adenhidayatuloh/Kelompok5_FinalProject3/pkg/errs"
)

type CategoryRepository interface {
	CreateCategory(category *entity.Category) (*entity.Category, errs.MessageErr)
	GetAllCategories() ([]entity.Category, errs.MessageErr)
	GetCategoryByID(id uint) (*entity.Category, errs.MessageErr)
	UpdateCategory(oldCategory *entity.Category, newCategory *entity.Category) (*entity.Category, errs.MessageErr)
	DeleteCategory(category *entity.Category) errs.MessageErr
}
