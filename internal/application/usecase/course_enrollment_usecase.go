package usecase

import (
	"capston-lms/internal/adapters/repository"
	"capston-lms/internal/entity"
)

type CourseEnrollmentUseCase struct {
	CourseEnrollmentRepo repository.CourseEnrollmentRepository
}

func (usecase CourseEnrollmentUseCase) GetStudents(id int) ([]entity.User, error) {
	courses, err := usecase.CourseEnrollmentRepo.GetAllStudents(id)
	return courses, err
}
func (usecase CourseEnrollmentUseCase) GetCourse(id int) ([]entity.Course, error) {
	courses, err := usecase.CourseEnrollmentRepo.GetCourse(id)
	return courses, err
}
