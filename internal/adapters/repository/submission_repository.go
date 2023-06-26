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
		Preload("Submission").
		First(&Submission, id)
	return Submission, result.Error
}

func (repo SubmissionRepository) CreateSubmission(Submission entity.Submission) error {
	result := repo.DB.Create(&Submission)
	return result.Error
}

func (repo SubmissionRepository) UpdateSubmission(id int, Submission entity.Submission) error {
	updates := make(map[string]interface{})
	// Add the desired columns and their values to the updates map
	if Submission.SubmissionSource != "" {
		updates["submission_source"] = Submission.SubmissionSource
	}
	if Submission.Status != "" {
		updates["status"] = Submission.Status
	}
	if Submission.Type != "" {
		updates["type"] = Submission.Type
	}
	if Submission.Notes != "" {
		updates["notes"] = Submission.Notes
	}
	if Submission.Score != "" {
		updates["score"] = Submission.Score
	}
	result := repo.DB.Model(&Submission).Where("id = ?", id).UpdateColumns(updates)

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
