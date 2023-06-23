package repository

import (
	"capston-lms/internal/entity"

	"gorm.io/gorm"
)

type RateCourseRepository struct {
	DB *gorm.DB
}

func (repo RateCourseRepository) CreateRateCourse(rateCourse entity.RateCourse) error {
	result := repo.DB.Create(&rateCourse)
	return result.Error
}

func (repo RateCourseRepository) UpdateStatusCourseEnrollment(courseID, userID int) error {
	err := repo.DB.Model(&entity.CourseEnrollment{}).
		Where("course_id = ? AND user_id = ?", courseID, userID).
		Update("status", "done").
		Error
	if err != nil {
		return err
	}

	return nil
}
