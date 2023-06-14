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
	result := repo.DB.Preload("Category").Preload("Class").Preload("Major").Find(&courses, "mentor_id = ?", id)
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

func (repo CourseRepository) GetCoursesByUserID(userID int) ([]entity.Course, error) {
	var courses []entity.Course
	result := repo.DB.Model(&entity.CourseEnrollment{}).
		Select("courses.*").
		Joins("JOIN courses ON course_enrollments.course_id = courses.ID").
		Where("course_enrollments.user_id = ?", userID).
		Find(&courses)
	return courses, result.Error
}

func (repo CourseRepository) CoursesInProgress(userID int) ([]entity.Course, error) {
	var courses []entity.Course
	result := repo.DB.Model(&entity.CourseEnrollment{}).
		Select("courses.*").
		Joins("JOIN courses ON course_enrollments.course_id = courses.ID").
		Where("course_enrollments.user_id = ?", userID).
		Where("course_enrollments.status = ?", "in_progress").
		Find(&courses)
	return courses, result.Error
}

func (repo CourseRepository) CoursesDone(userID int) ([]entity.Course, error) {
	var courses []entity.Course
	result := repo.DB.Model(&entity.CourseEnrollment{}).
		Select("courses.*").
		Joins("JOIN courses ON course_enrollments.course_id = courses.ID").
		Where("course_enrollments.user_id = ?", userID).
		Where("course_enrollments.status = ?", "selesai").
		Find(&courses)
	return courses, result.Error
}

func (repo CourseRepository) GetAllModules() ([]entity.Module, error) {
	var modules []entity.Module
	result := repo.DB.Find(&modules)
	return modules, result.Error
}

func (repo CourseRepository) GetModule(id int) (entity.Module, error) {
	var module entity.Module
	result := repo.DB.Preload("Section").
		Preload("Tasks").
		Preload("Attachment").
		Preload("Submission").
		Preload("Submission.User").
		Preload("Section.Course").
		First(&module, id)
	return module, result.Error
}
