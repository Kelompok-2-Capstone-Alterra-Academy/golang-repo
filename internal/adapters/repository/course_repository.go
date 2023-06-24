package repository

import (
	"capston-lms/internal/entity"

	"gorm.io/gorm"
)

type CourseRepository struct {
	DB *gorm.DB
}

func (repo CourseRepository) GetAllCourses(mentorId int) ([]entity.Course, error) {
	var courses []entity.Course
	result := repo.DB.Where("mentor_id = ?", mentorId).Preload("Category").
		Preload("Section", func(db *gorm.DB) *gorm.DB {
			return db.Omit("course") // Menyembunyikan relasi "Course" pada preload "Section"
		}).
		Preload("Class").Preload("Major").Find(&courses)
	return courses, result.Error
}
func (repo CourseRepository) GetAllCourseStudents() ([]entity.Course, error) {
	var courses []entity.Course
	result := repo.DB.Where("status = ?", "publish").Preload("Category").
		Preload("Section", func(db *gorm.DB) *gorm.DB {
			return db.Omit("course") // Menyembunyikan relasi "Course" pada preload "Section"
		}).
		Preload("Class").Preload("Major").Find(&courses)
	// Mendapatkan jumlah siswa yang mengikuti setiap kursus
	for i := range courses {
		var count int64
		repo.DB.Model(&entity.CourseEnrollment{}).Where("course_id = ?", courses[i].ID).Count(&count)
		courses[i].NumStudents = int(count)
	}
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

func (repo CourseRepository) CreateCourse(course *entity.Course) error {
	result := repo.DB.Preload("Category").Preload("Class").Preload("Major").Create(&course)
	return result.Error
}

func (repo CourseRepository) UpdateCourse(id int, course entity.Course) error {
	result := repo.DB.Model(&course).Where("id = ?", id).UpdateColumns(course)
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
		Preload("Class").
		Preload("Major").
		Find(&courses)
	return courses, result.Error
}

func (repo CourseRepository) CoursesDone(userID int) ([]entity.Course, error) {
	var courses []entity.Course
	result := repo.DB.Model(&entity.CourseEnrollment{}).
		Select("courses.*").
		Joins("JOIN courses ON course_enrollments.course_id = courses.ID").
		Where("course_enrollments.user_id = ?", userID).
		Preload("Class").
		Preload("Major").
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
func (repo CourseRepository) GetAllCoursesWithSectionCount(courseId int) ([]entity.CourseWithSectionCount, error) {
	var courses []entity.CourseWithSectionCount

	result := repo.DB.Table("courses").
		Select("courses.*, COUNT(sections.id) AS section_count").
		Where("courses.id = ?", courseId).
		Joins("LEFT JOIN sections ON courses.id = sections.course_id").
		Group("courses.id").
		Scan(&courses)

	if result.Error != nil {
		return nil, result.Error
	}

	return courses, nil
}

func (repo CourseRepository) GetCourseSection(id int) (entity.Course, error) {
	var courses entity.Course
	result := repo.DB.Preload("Section").Preload("Section.Module").Preload("Section.Module.Attachment").First(&courses, id)
	return courses, result.Error
}

func (repo CourseRepository) GetStudentsByCourseID(courseID int) ([]entity.User, error) {
	var users []entity.User
	result := repo.DB.Model(&entity.CourseEnrollment{}).
		Select("users.*").
		Joins("JOIN users ON course_enrollments.user_id = users.ID").
		Where("course_enrollments.course_id = ?", courseID).
		Where("course_enrollments.status <> ?", "deactive").
		Find(&users)
	return users, result.Error
}

func (repo CourseRepository) GetAllCoursesWithSectionAndStudentCount() ([]entity.CourseWithSectionAndStudentCount, error) {
	var courses []entity.CourseWithSectionAndStudentCount

	result := repo.DB.Table("courses").
		Select("courses.*, COUNT(DISTINCT sections.id) AS section_count, COUNT(DISTINCT CASE WHEN users.role = 'students' THEN users.id END) AS student_count").
		Joins("LEFT JOIN sections ON courses.id = sections.course_id").
		Joins("LEFT JOIN course_enrollments ON courses.id = course_enrollments.course_id").
		Joins("LEFT JOIN users ON course_enrollments.user_id = users.id").
		Group("courses.id").
		Scan(&courses)

	if result.Error != nil {
		return nil, result.Error
	}

	return courses, nil
}
