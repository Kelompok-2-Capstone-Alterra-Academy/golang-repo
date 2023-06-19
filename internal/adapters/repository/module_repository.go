package repository

import (
	"capston-lms/internal/entity"

	"gorm.io/gorm"
)

type ModuleRepository struct {
	DB *gorm.DB
}

func (repo ModuleRepository) GetAllModules() ([]entity.Module, error) {
	var modules []entity.Module
	result := repo.DB.Preload("Section").
		Preload("Tasks").
		Preload("Attachment").
		Preload("Submission").
		Preload("Submission.User").
		Find(&modules)
	return modules, result.Error
}

func (repo ModuleRepository) GetModule(id int) (entity.Module, error) {
	var module entity.Module
	result := repo.DB.Preload("Section").
		Preload("Tasks").
		Preload("Attachment").
		Preload("Submission").
		Preload("Submission.User").
		First(&module, id)
	return module, result.Error
}

func (repo ModuleRepository) CreateModule(module *entity.Module) error {
	result := repo.DB.Create(&module)
	return result.Error
}

func (repo ModuleRepository) UpdateModule(id int, module entity.Module) error {

	result := repo.DB.Model(&module).Where("ids = ?", id).UpdateColumns(module)
	return result.Error
}

func (repo ModuleRepository) DeleteModule(id int) error {
	result := repo.DB.Delete(&entity.Module{}, id)
	return result.Error
}

func (repo ModuleRepository) FindModule(id int) error {
	result := repo.DB.First(&entity.Module{}, id)
	return result.Error
}
