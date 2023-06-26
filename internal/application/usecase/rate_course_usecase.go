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

func (usecase RateCourseUseCase) UpdateStatusCourseEnrollment(courseID, userID int) error {
	err := usecase.Repo.UpdateStatusCourseEnrollment(courseID, userID)
	return err
}
