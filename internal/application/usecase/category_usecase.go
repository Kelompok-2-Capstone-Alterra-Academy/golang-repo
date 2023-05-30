package usecase

import (
	"capston-lms/internal/adapters/repository"
	"capston-lms/internal/entity"
)

type CategoryUseCase struct {
	Repo repository.CategoryRepository
}

func (usecase CategoryUseCase) GetAllCategories() ([]entity.Category, error) {
	categoryes, err := usecase.Repo.GetAllCategories()
	return categoryes, err
}

func (usecase CategoryUseCase) GetCategory(id int) (entity.Category, error) {
	category, err := usecase.Repo.GetCategory(id)
	return category, err
}

func (usecase CategoryUseCase) CreateCategory(category entity.Category) error {
	err := usecase.Repo.CreateCategory(category)
	return err
}

func (usecase CategoryUseCase) UpdateCategory(id int, category entity.Category) error {
	err := usecase.Repo.UpdateCategory(id, category)
	return err
}

func (usecase CategoryUseCase) DeleteCategory(id int) error {
	err := usecase.Repo.DeleteCategory(id)
	return err
}
func (usecase CategoryUseCase) FindCategory(id int) error {
	err := usecase.Repo.FindCategory(id)
	return err
}
