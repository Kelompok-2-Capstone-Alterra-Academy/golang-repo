package usecase

import (
	"capston-lms/internal/adapters/repository"
	"capston-lms/internal/entity"
)

type RateCourseUseCase struct {
	Repo repository.RateCourseRepository
}

func (usecase RateCourseUseCase) CreateRateCourse(rateCourse entity.RateCourse) error {
	err := usecase.Repo.CreateRateCourse(rateCourse)
	return err
}
