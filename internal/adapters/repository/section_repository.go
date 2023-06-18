package repository

import (
	"capston-lms/internal/entity"

	"gorm.io/gorm"
)

type SectionRepository struct {
	DB *gorm.DB
}

func (repo SectionRepository) GetAllSections() ([]entity.Section, error) {
	var sections []entity.Section
	result := repo.DB.Find(&sections)
	return sections, result.Error
}

func (repo SectionRepository) GetSection(id int) (entity.Section, error) {
	var sections entity.Section
	result := repo.DB.First(&sections, id)
	return sections, result.Error
}

func (repo SectionRepository) CreateSection(section *entity.Section) error {
	result := repo.DB.Create(&section)
	return result.Error
}

func (repo SectionRepository) UpdateSection(id int, section *entity.Section) error {
	result := repo.DB.Model(&section).Where("id = ?", id).Updates(&section)
	return result.Error
}

func (repo SectionRepository) DeleteSection(id int) error {
	result := repo.DB.Delete(&entity.Section{}, id)
	return result.Error
}

func (repo SectionRepository) FindSection(id int) error {
	result := repo.DB.First(&entity.Section{}, id)
	return result.Error
}
func (repo SectionRepository) GetAllSectionsByCourse(course_id int) ([]entity.Section, error) {
	var sections []entity.Section
	result := repo.DB.Preload("Module").
		Preload("Module.Attachment").
		Preload("Course").
		Where("course_id = ?", course_id).
		Find(&sections)
	return sections, result.Error
}
