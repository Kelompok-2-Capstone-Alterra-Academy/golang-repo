package repository

import (
	"capston-lms/internal/entity"

	"gorm.io/gorm"
)

type MajorRepository struct {
	DB *gorm.DB
}

func (repo MajorRepository) GetAllMajors() ([]entity.Major, error) {
	var majors []entity.Major
	result := repo.DB.Find(&majors)
	return majors, result.Error
}

func (repo MajorRepository) GetMajor(id int) (entity.Major, error) {
	var majors entity.Major
	result := repo.DB.First(&majors, id)
	return majors, result.Error
}

func (repo MajorRepository) CreateMajor(major entity.Major) error {
	result := repo.DB.Create(&major)
	return result.Error
}

func (repo MajorRepository) UpdateMajor(id int, major entity.Major) error {
	result := repo.DB.Model(&major).Where("id = ?", id).Updates(&major)
	return result.Error
}

func (repo MajorRepository) DeleteMajor(id int) error {
	result := repo.DB.Delete(&entity.Major{}, id)
	return result.Error
}

func (repo MajorRepository) FindMajor(id int) error {
	result := repo.DB.First(&entity.Major{}, id)
	return result.Error
}
