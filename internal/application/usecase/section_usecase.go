package usecase

import (
	"capston-lms/internal/adapters/repository"
	"capston-lms/internal/entity"
)

type SectionUseCase struct {
	Repo repository.SectionRepository
}

func (usecase SectionUseCase) GetAllSections() ([]entity.Section, error) {
	sections, err := usecase.Repo.GetAllSections()
	return sections, err
}

func (usecase SectionUseCase) GetSection(id int) (entity.Section, error) {
	section, err := usecase.Repo.GetSection(id)
	return section, err
}

func (usecase SectionUseCase) CreateSection(section *entity.Section) error {
	err := usecase.Repo.CreateSection(section)
	return err
}

func (usecase SectionUseCase) UpdateSection(id int, section *entity.Section) error {
	err := usecase.Repo.UpdateSection(id, section)
	return err
}

func (usecase SectionUseCase) DeleteSection(id int) error {
	err := usecase.Repo.DeleteSection(id)
	return err
}

func (usecase SectionUseCase) FindSection(id int) error {
	err := usecase.Repo.FindSection(id)
	return err
}

func (usecase SectionUseCase) GetAllSectionsByCourse(course int) ([]entity.Section, error) {
	sections, err := usecase.Repo.GetAllSectionsByCourse(course)
	return sections, err
}
