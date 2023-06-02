package repository

import (
	"capston-lms/internal/entity"

	"gorm.io/gorm"
)

type CourseEnrollmentRepository struct {
	DB *gorm.DB
}

func (repo CourseEnrollmentRepository) GetAllStudents(course_id int) ([]entity.CourseEnrollment, error) {
	var courses []entity.CourseEnrollment
	result := repo.DB.Preload("User").Where("course_enrollments.course_id = ?", course_id).Find(&courses)

	return courses, result.Error
}

func (repo CourseEnrollmentRepository) GetCourse(userID int) ([]entity.Course, error) {
	var courses []entity.Course
	result := repo.DB.Where("courses.mentor_id = ?", userID).Find(&courses)

	return courses, result.Error
}
