package repository

import (
	"capston-lms/internal/entity"

	"gorm.io/gorm"
)

type AttachmentRepository struct {
	DB *gorm.DB
}

func (repo AttachmentRepository) GetAttachmentByCourseSectionIDAndAttachmentID(courseSectionID string, attachmentID string) (*entity.Attachment, error) {
	attachment := &entity.Attachment{}
	result := repo.DB.Preload("Mentor").
		Where("id = ? AND course_section_id = ?", attachmentID, courseSectionID).
		First(attachment)
	if result.Error != nil {
		return nil, result.Error
	}

	return attachment, nil
}
