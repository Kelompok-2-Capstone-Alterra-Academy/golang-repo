package usecase

import (
	"capston-lms/internal/adapters/repository"
	"capston-lms/internal/entity"
)

type ClassUseCase struct {
	Repo repository.ClassRepository
}

func (usecase ClassUseCase) GetAllClasses() ([]entity.Class, error) {
	classes, err := usecase.Repo.GetAllClasses()
	return classes, err
}

func (usecase ClassUseCase) GetClass(id int) (entity.Class, error) {
	class, err := usecase.Repo.GetClass(id)
	return class, err
}

func (usecase ClassUseCase) CreateClass(class entity.Class) error {
	err := usecase.Repo.CreateClass(class)
	return err
}

func (usecase ClassUseCase) UpdateClass(id int, class entity.Class) error {
	err := usecase.Repo.UpdateClass(id, class)
	return err
}

func (usecase ClassUseCase) DeleteClass(id int) error {
	err := usecase.Repo.DeleteClass(id)
	return err
}
func (usecase ClassUseCase) FindClass(id int) error {
	err := usecase.Repo.FindClass(id)
	return err
}
