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
	result := repo.DB.Preload("Category").Preload("Class").Preload("Major").Find(&courses)
	return courses, result.Error
}

func (repo CourseRepository) GetCourse(id int) (entity.Course, error) {
	var courses entity.Course
	result := repo.DB.Preload("Category").Preload("Class").Preload("Major").First(&courses, id)
	return courses, result.Error
}
func (repo CourseRepository) GetCourseByMentorId(id int) (entity.Course, error) {
	var courses entity.Course
	result := repo.DB.Preload("Category").Preload("Class").Preload("Major").Find(&courses,"mentor_id = ?", id)
	return courses, result.Error
}

func (repo CourseRepository) CreateCourse(course entity.Course) error {
	result := repo.DB.Preload("Category").Preload("Class").Preload("Major").Create(&course)
	return result.Error
}

func (repo CourseRepository) UpdateCourse(id int, course entity.Course) error {
	result := repo.DB.Preload("Category").Preload("Class").Preload("Major").Model(&course).Where("id = ?", id).Updates(&course)
	return result.Error
}

func (repo CourseRepository) DeleteCourse(id int) error {
	result := repo.DB.Delete(&entity.Course{}, id)
	return result.Error
}
func (repo CourseRepository) FindCourse(id int) error {
	result := repo.DB.Preload("Category").Preload("Class").Preload("Major").First(&entity.Course{}, id)
	return result.Error
}

func (repo CourseRepository) GetAllCoursesSortedByCompletion(ascending bool) ([]entity.Course, error) {
	var courses []entity.Course
	order := "ASC"
	if !ascending {
		order = "DESC"
	}
	result := repo.DB.Preload("Category").Preload("Class").Preload("Major").Order("completion " + order).Find(&courses)
	return courses, result.Error
}

func (repo CourseRepository) GetAllCoursesSortedByNewness(ascending bool) ([]entity.Course, error) {
	var courses []entity.Course
	order := "ASC"
	if !ascending {
		order = "DESC"
	}
	result := repo.DB.Preload("Category").Preload("Class").Preload("Major").Order("created_at " + order).Find(&courses)
	return courses, result.Error
}

func (repo CourseRepository) GetAllCoursesSortedByHighLevel(ascending bool) ([]entity.Course, error) {
	var courses []entity.Course
	order := "ASC"
	if !ascending {
		order = "DESC"
	}
	result := repo.DB.Preload("Category").Preload("Class").Preload("Major").Order("level " + order).Find(&courses)
	return courses, result.Error
}

func (repo CourseRepository) GetAllCoursesSortedByLowLevel(ascending bool) ([]entity.Course, error) {
	var courses []entity.Course
	order := "ASC"
	if !ascending {
		order = "DESC"
	}
	result := repo.DB.Preload("Category").Preload("Class").Preload("Major").Order("level " + order).Find(&courses)
	return courses, result.Error
}