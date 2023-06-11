package usecase

import (
	"capston-lms/internal/adapters/repository"
	"capston-lms/internal/entity"
)

type MajorUseCase struct {
	Repo repository.MajorRepository
}

func (usecase MajorUseCase) GetAllMajors() ([]entity.Major, error) {
	majores, err := usecase.Repo.GetAllMajors()
	return majores, err
}

func (usecase MajorUseCase) GetMajor(id int) (entity.Major, error) {
	major, err := usecase.Repo.GetMajor(id)
	return major, err
}

func (usecase MajorUseCase) CreateMajor(major entity.Major) error {
	err := usecase.Repo.CreateMajor(major)
	return err
}

func (usecase MajorUseCase) UpdateMajor(id int, major entity.Major) error {
	err := usecase.Repo.UpdateMajor(id, major)
	return err
}

func (usecase MajorUseCase) DeleteMajor(id int) error {
	err := usecase.Repo.DeleteMajor(id)
	return err
}

func (usecase MajorUseCase) FindMajor(id int) error {
	err := usecase.Repo.FindMajor(id)
	return err
}
