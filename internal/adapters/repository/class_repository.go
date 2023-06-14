package repository

import (
	"capston-lms/internal/entity"

	"gorm.io/gorm"
)

type ClassRepository struct {
	DB *gorm.DB
}

func (repo ClassRepository) GetAllClasses() ([]entity.Class, error) {
	var classes []entity.Class
	result := repo.DB.Find(&classes)
	return classes, result.Error
}

func (repo ClassRepository) GetClass(id int) (entity.Class, error) {
	var classes entity.Class
	result := repo.DB.First(&classes, id)
	return classes, result.Error
}

func (repo ClassRepository) CreateClass(class entity.Class) error {
	result := repo.DB.Create(&class)
	return result.Error
}

func (repo ClassRepository) UpdateClass(id int, class entity.Class) error {
	result := repo.DB.Model(&class).Where("id = ?", id).Updates(&class)
	return result.Error
}

func (repo ClassRepository) DeleteClass(id int) error {
	result := repo.DB.Delete(&entity.Class{}, id)
	return result.Error
}
func (repo ClassRepository) FindClass(id int) error {
	result := repo.DB.First(&entity.Class{}, id)
	return result.Error
}
