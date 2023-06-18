package usecase

import (
	"capston-lms/internal/adapters/repository"
	"capston-lms/internal/entity"
)

type ModuleUseCase struct {
	Repo repository.ModuleRepository
}

func (usecase ModuleUseCase) GetAllModules() ([]entity.Module, error) {
	modules, err := usecase.Repo.GetAllModules()
	return modules, err
}

func (usecase ModuleUseCase) GetModule(id int) (entity.Module, error) {
	module, err := usecase.Repo.GetModule(id)
	return module, err
}

func (usecase ModuleUseCase) CreateModule(module *entity.Module) error {
	err := usecase.Repo.CreateModule(module)
	return err
}

func (usecase ModuleUseCase) UpdateModule(id int, module *entity.Module) error {
	err := usecase.Repo.UpdateModule(id, module)
	return err
}

func (usecase ModuleUseCase) DeleteModule(id int) error {
	err := usecase.Repo.DeleteModule(id)
	return err
}

func (usecase ModuleUseCase) FindModule(id int) error {
	err := usecase.Repo.FindModule(id)
	return err
}
