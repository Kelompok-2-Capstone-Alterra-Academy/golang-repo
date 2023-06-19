package repository

import (
	"capston-lms/internal/entity"

	"gorm.io/gorm"
)

type CourseEnrollmentRepository struct {
	DB *gorm.DB
}

func (repo CourseEnrollmentRepository) GetAllStudents(course_id int) ([]entity.User, error) {

	var users []entity.User
	result := repo.DB.Model(&entity.CourseEnrollment{}).
		Select("users.*").
		Joins("JOIN users ON course_enrollments.user_id = users.ID").
		Where("course_enrollments.course_id = ?", course_id).
		Find(&users)
	return users, result.Error
}

func (repo CourseEnrollmentRepository) GetCourse(userID int) ([]entity.Course, error) {
	var courses []entity.Course
	result := repo.DB.Where("courses.mentor_id = ?", userID).Find(&courses)

	return courses, result.Error
}

func (repo CourseEnrollmentRepository) DeleteCourseEnrollment(userId int, courseId int, course entity.CourseEnrollment) error {
	result := repo.DB.Model(&entity.CourseEnrollment{}).Where("user_id = ? AND course_id = ?", userId, courseId).Update("status", "deactive")
	return result.Error
}
