package repository

import (
	"capston-lms/internal/entity"

	"gorm.io/gorm"
)

type EducationNewsRepository struct {
	DB *gorm.DB
}

func (repo EducationNewsRepository) GetAllEducationNewses() ([]entity.EducationNews, error) {
	var education_newses []entity.EducationNews
	result := repo.DB.Find(&education_newses)
	return education_newses, result.Error
}

func (repo EducationNewsRepository) GetEducationNews(id int) (entity.EducationNews, error) {
	var education_newses entity.EducationNews
	result := repo.DB.First(&education_newses, id)
	return education_newses, result.Error
}

func (repo EducationNewsRepository) CreateEducationNews(education_news entity.EducationNews) error {
	result := repo.DB.Create(&education_news)
	return result.Error
}

func (repo EducationNewsRepository) UpdateEducationNews(id int, education_news entity.EducationNews) error {
	result := repo.DB.Model(&education_news).Where("id = ?", id).Updates(&education_news)
	return result.Error
}

func (repo EducationNewsRepository) DeleteEducationNews(id int) error {
	result := repo.DB.Delete(&entity.EducationNews{}, id)
	return result.Error
}

func (repo EducationNewsRepository) FindEducationNews(id int) error {
	result := repo.DB.First(&entity.EducationNews{}, id)
	return result.Error
}
