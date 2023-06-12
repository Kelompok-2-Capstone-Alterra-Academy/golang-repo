package repository

import (
	"capston-lms/internal/entity"

	"gorm.io/gorm"
)

type SubmissionRepository struct {
	DB *gorm.DB
}

func (repo SubmissionRepository) GetAllSubmissions() ([]entity.Submission, error) {
	var Submissions []entity.Submission
	result := repo.DB.Preload("User").
		Find(&Submissions)
	return Submissions, result.Error
}

func (repo SubmissionRepository) GetSubmission(id int) (entity.Submission, error) {
	var Submission entity.Submission
	result := repo.DB.Preload("User").
		Preload("Module").
		First(&Submission, id)
	return Submission, result.Error
}

func (repo SubmissionRepository) CreateSubmission(Submission entity.Submission) error {
	result := repo.DB.Create(&Submission)
	return result.Error
}

func (repo SubmissionRepository) UpdateSubmission(id int, Submission entity.Submission) error {

	result := repo.DB.Model(&Submission).Where("ids = ?", id).Updates(&Submission)
	return result.Error
}

func (repo SubmissionRepository) DeleteSubmission(id int) error {
	result := repo.DB.Delete(&entity.Submission{}, id)
	return result.Error
}

func (repo SubmissionRepository) FindSubmission(id int) error {
	result := repo.DB.First(&entity.Submission{}, id)
	return result.Error
}
