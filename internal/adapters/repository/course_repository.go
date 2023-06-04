package repository

import (
	"capston-lms/internal/entity"

	"gorm.io/gorm"
)

type CourseRepository struct {
	DB *gorm.DB
}

func (repo CourseRepository) GetAllCourses() ([]entity.Course, error) {
	var courses []entity.Course
	result := repo.DB.Find(&courses)
	return courses, result.Error
}

func (repo CourseRepository) GetCourse(id int) (entity.Course, error) {
	var courses entity.Course
	result := repo.DB.First(&courses, id)
	return courses, result.Error
}

func (repo CourseRepository) CreateCourse(course entity.Course) error {
	result := repo.DB.Preload("Category").Preload("Class").Preload("Major").Create(&course)
	return result.Error
}

func (repo CourseRepository) UpdateCourse(id int, course entity.Course) error {
	result := repo.DB.Model(&course).Where("id = ?", id).Updates(&course)
	return result.Error
}

func (repo CourseRepository) DeleteCourse(id int) error {
	result := repo.DB.Delete(&entity.Course{}, id)
	return result.Error
}
func (repo CourseRepository) FindCourse(id int) error {
	result := repo.DB.First(&entity.Course{}, id)
	return result.Error
}

func (repo CourseRepository) GetCourseByIDAndUserID(courseID string, userID string) (*entity.Course, error) {
	course := &entity.Course{}
	result := repo.DB.Preload("Category").
		Preload("Class").
		Preload("Mentor").
		Preload("Major").
		Where("id = ? AND mentor_id = ?", courseID, userID).
		First(course)
	if result.Error != nil {
		return nil, result.Error
	}

	return course, nil
}

func (repo CourseRepository) GetCourseByID(courseID string) (*entity.Course, error) {
	course := &entity.Course{}
	result := repo.DB.Preload("Category").
		Preload("Class").
		Preload("Mentor").
		Preload("Major").
		Where("id = ?", courseID).
		First(course)
	if result.Error != nil {
		return nil, result.Error
	}

	return course, nil
}
