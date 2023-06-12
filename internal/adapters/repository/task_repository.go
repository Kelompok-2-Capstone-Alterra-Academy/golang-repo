package repository

import (
	"capston-lms/internal/entity"

	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func (repo TaskRepository) GetAllTasks() ([]entity.Task, error) {
	var Tasks []entity.Task
	result := repo.DB.Find(&Tasks)
	return Tasks, result.Error
}

func (repo TaskRepository) GetTask(id int) (entity.Task, error) {
	var Tasks entity.Task
	result := repo.DB.Where("module_id = ?", id).First(&Tasks)
	return Tasks, result.Error
}

func (repo TaskRepository) CreateTask(Task entity.Task) error {
	result := repo.DB.Create(&Task)
	return result.Error
}

func (repo TaskRepository) UpdateTask(id int, Task entity.Task) error {
	result := repo.DB.Model(&Task).Where("id = ?", id).Updates(&Task)
	return result.Error
}

func (repo TaskRepository) DeleteTask(id int) error {
	result := repo.DB.Delete(&entity.Task{}, id)
	return result.Error
}

func (repo TaskRepository) FindTask(id int) error {
	result := repo.DB.First(&entity.Task{}, id)
	return result.Error
}
