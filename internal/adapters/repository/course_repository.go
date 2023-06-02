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
