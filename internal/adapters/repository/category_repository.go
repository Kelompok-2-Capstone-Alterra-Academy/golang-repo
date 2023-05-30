package repository

import (
	"capston-lms/internal/entity"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func (repo CategoryRepository) GetAllCategories() ([]entity.Category, error) {
	var categories []entity.Category
	result := repo.DB.Find(&categories)
	return categories, result.Error
}

func (repo CategoryRepository) GetCategory(id int) (entity.Category, error) {
	var categories entity.Category
	result := repo.DB.First(&categories, id)
	return categories, result.Error
}

func (repo CategoryRepository) CreateCategory(category entity.Category) error {
	result := repo.DB.Create(&category)
	return result.Error
}

func (repo CategoryRepository) UpdateCategory(id int, category entity.Category) error {
	result := repo.DB.Model(&category).Where("id = ?", id).Updates(&category)
	return result.Error
}

func (repo CategoryRepository) DeleteCategory(id int) error {
	result := repo.DB.Delete(&entity.Category{}, id)
	return result.Error
}
func (repo CategoryRepository) FindCategory(id int) error {
	result := repo.DB.First(&entity.Category{}, id)
	return result.Error
}

