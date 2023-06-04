package usecase

import (
	"capston-lms/internal/adapters/repository"
	"capston-lms/internal/entity"
)

type CourseUseCase struct {
	Repo repository.CourseRepository
}

func (usecase CourseUseCase) GetAllCourses() ([]entity.Course, error) {
	courses, err := usecase.Repo.GetAllCourses()
	return courses, err
}

func (usecase CourseUseCase) GetCourse(id int) (entity.Course, error) {
	course, err := usecase.Repo.GetCourse(id)
	return course, err
}

func (usecase CourseUseCase) CreateCourse(course entity.Course) error {
	err := usecase.Repo.CreateCourse(course)
	return err
}

func (usecase CourseUseCase) UpdateCourse(id int, course entity.Course) error {
	err := usecase.Repo.UpdateCourse(id, course)
	return err
}

func (usecase CourseUseCase) DeleteCourse(id int) error {
	err := usecase.Repo.DeleteCourse(id)
	return err
}
func (usecase CourseUseCase) FindCourse(id int) error {
	err := usecase.Repo.FindCourse(id)
	return err
}

func (usecase CourseUseCase) GetCourseByIDAndUserID(courseID string, userID string) (*entity.Course, error) {
	course, err := usecase.Repo.GetCourseByIDAndUserID(courseID, userID)
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (usecase CourseUseCase) GetCourseByID(courseID string) (*entity.Course, error) {
	course, err := usecase.Repo.GetCourseByID(courseID)
	if err != nil {
		return nil, err
	}

	return course, nil
}
