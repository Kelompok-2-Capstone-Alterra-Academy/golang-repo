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
func (usecase CourseUseCase) GetCourseByMentorId(id int) (entity.Course, error) {
	course, err := usecase.Repo.GetCourseByMentorId(id)
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

func (usecase CourseUseCase) GetAllCoursesSortedByCompletion(ascending bool) ([]entity.Course, error) {
	courses, err := usecase.Repo.GetAllCoursesSortedByCompletion(ascending)
	return courses, err
}

func (usecase CourseUseCase) GetAllCoursesSortedByNewness(ascending bool) ([]entity.Course, error) {
	courses, err := usecase.Repo.GetAllCoursesSortedByNewness(ascending)
	return courses, err
}

func (usecase CourseUseCase) GetAllCoursesSortedByHighLevel(ascending bool) ([]entity.Course, error) {
	courses, err := usecase.Repo.GetAllCoursesSortedByHighLevel(ascending)
	return courses, err
}

func (usecase CourseUseCase) GetAllCoursesSortedByLowLevel(ascending bool) ([]entity.Course, error) {
	courses, err := usecase.Repo.GetAllCoursesSortedByLowLevel(ascending)
	return courses, err
}



// func (usecase CourseUseCase) GetAllCoursesSortedByHighLevel(ascending bool) ([]entity.Course, error) {
// 	courses, err := usecase.Repo.GetAllCoursesSortedByHighLevel(ascending)
// 	return courses, err
// }

// func (usecase CourseUseCase) GetAllCoursesSortedByLowLevel(ascending bool) ([]entity.Course, error) {
// 	courses, err := usecase.Repo.GetAllCoursesSortedByLowLevel(ascending)
// 	return courses, err
// }

// func (repo CourseUseCase) GetAllCoursesSortedByHighLevel(ascending bool) ([]entity.Course, error) {
// 	var courses []entity.Course
// 	order := "ASC"
// 	if !ascending {
// 		order = "DESC"
// 	}
// 	result := repo.DB.Preload("Category").Preload("Class").Preload("Major").Order("level " + order).Find(&courses)
// 	return courses, result.Error
// }
