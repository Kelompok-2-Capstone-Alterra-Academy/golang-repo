package usecase

import (
	"capston-lms/internal/adapters/repository"
	"capston-lms/internal/entity"
)

type EducationNewsUseCase struct {
	Repo repository.EducationNewsRepository
}

func (usecase EducationNewsUseCase) GetAllEducationNewses() ([]entity.EducationNews, error) {
	education_newses, err := usecase.Repo.GetAllEducationNewses()
	return education_newses, err
}

func (usecase EducationNewsUseCase) GetEducationNews(id int) (entity.EducationNews, error) {
	education_news, err := usecase.Repo.GetEducationNews(id)
	return education_news, err
}

func (usecase EducationNewsUseCase) CreateEducationNews(education_news entity.EducationNews) error {
	err := usecase.Repo.CreateEducationNews(education_news)
	return err
}

func (usecase EducationNewsUseCase) UpdateEducationNews(id int, education_news entity.EducationNews) error {
	err := usecase.Repo.UpdateEducationNews(id, education_news)
	return err
}

func (usecase EducationNewsUseCase) DeleteEducationNews(id int) error {
	err := usecase.Repo.DeleteEducationNews(id)
	return err
}

func (usecase EducationNewsUseCase) FindEducationNews(id int) error {
	err := usecase.Repo.FindEducationNews(id)
	return err
}
