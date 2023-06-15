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
