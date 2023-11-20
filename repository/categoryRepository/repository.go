package categoryrepository

import (
	"github.com/adenhidayatuloh/glng_ks08_Kelompok5_final_Project_3/entity"
	"github.com/adenhidayatuloh/glng_ks08_Kelompok5_final_Project_3/pkg/errs"
)

type CategoryRepository interface {
	CreateCategory(category *entity.Category) (*entity.Category, errs.MessageErr)
	GetAllCategories() ([]entity.Category, errs.MessageErr)
	GetCategoryByID(id uint) (*entity.Category, errs.MessageErr)
	UpdateCategory(oldCategory *entity.Category, newCategory *entity.Category) (*entity.Category, errs.MessageErr)
	DeleteCategory(category *entity.Category) errs.MessageErr
}
