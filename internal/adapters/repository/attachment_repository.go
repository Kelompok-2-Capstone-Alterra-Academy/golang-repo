package repository

import (
	"capston-lms/internal/entity"

	"gorm.io/gorm"
)

type AttachmentRepository struct {
	DB *gorm.DB
}

func (repo AttachmentRepository) GetAllAttachments(folder int) ([]entity.Attachment, error) {
	var Attachments []entity.Attachment
	result := repo.DB.Where("folder_id = ?", folder).Find(&Attachments)
	if result.Error != nil {
		return nil, result.Error
	}
	return Attachments, nil
}
func (repo AttachmentRepository) GetQuiz() ([]entity.Attachment, error) {
	var Attachments []entity.Attachment
	var quiz = "quiz"
	result := repo.DB.Where("type = ?", quiz).Find(&Attachments)
	if result.Error != nil {
		return nil, result.Error
	}
	return Attachments, nil
}
func (repo AttachmentRepository) GetAttachment(id int) (entity.Attachment, error) {
	var Attachments entity.Attachment
	result := repo.DB.First(&Attachments, id)
	return Attachments, result.Error
}

func (repo AttachmentRepository) CreateAttachment(Attachment *entity.Attachment) error {
	result := repo.DB.Create(&Attachment)
	return result.Error
}

func (repo AttachmentRepository) UpdateAttachment(id int, attachment entity.Attachment) error {
	updates := make(map[string]interface{})
	// Add the desired columns and their values to the updates map
	if attachment.Status != "" {
		updates["status"] = attachment.Status
	}
	if attachment.Type != "" {
		updates["type"] = attachment.Type
	}
	if attachment.AttachmentName != "" {
		updates["attachment_names"] = attachment.AttachmentName
	}
	if attachment.Description != "" {
		updates["description"] = attachment.Description
	}

	if attachment.AttachmentSource != "" {
		updates["attachment_source"] = attachment.AttachmentSource
	}

	if attachment.FolderId != nil && *attachment.FolderId != "" {
		updates["folder_id"] = attachment.FolderId
	}

	result := repo.DB.Model(&attachment).Where("id = ?", id).UpdateColumns(updates)
	return result.Error
}

func (repo AttachmentRepository) DeleteAttachment(id int) error {
	result := repo.DB.Delete(&entity.Attachment{}, id)
	return result.Error
}

func (repo AttachmentRepository) FindAttachment(id int) error {
	result := repo.DB.First(&entity.Attachment{}, id)
	return result.Error
}

func (repo AttachmentRepository) GetVideoAttachments() ([]entity.Attachment, error) {
	var attachments []entity.Attachment

	// Query the attachments with type "video"
	if err := repo.DB.Where("type = ?", "video").Find(&attachments).Error; err != nil {
		return nil, err
	}

	// Return the attachments
	return attachments, nil
}

func (repo AttachmentRepository) GetVideoAttachmentByID(id int) (entity.Attachment, error) {
	var attachment entity.Attachment
	result := repo.DB.First(&attachment, id)
	return attachment, result.Error
}

func (repo AttachmentRepository) GetQuizAttachments() ([]entity.Attachment, error) {
	var attachments []entity.Attachment

	// Query the attachments with type "quiz"
	if err := repo.DB.Where("type = ?", "quiz").Find(&attachments).Error; err != nil {
		// Handle the error
		return nil, err
	}

	// Return the attachments
	return attachments, nil
}

func (repo AttachmentRepository) GetQuizAttachmentByID(id int) (entity.Attachment, error) {
	var attachment entity.Attachment
	result := repo.DB.First(&attachment, id)
	return attachment, result.Error
}

func (repo AttachmentRepository) GetMateriAttachments() ([]entity.Attachment, error) {
	var attachments []entity.Attachment

	// Query the attachments with type "materi"
	if err := repo.DB.Where("type = ?", "materi").Find(&attachments).Error; err != nil {
		// Handle the error
		return nil, err
	}

	// Return the attachments
	return attachments, nil
}

func (repo AttachmentRepository) GetMateriAttachmentByID(id int) (entity.Attachment, error) {
	var attachment entity.Attachment
	result := repo.DB.First(&attachment, id)
	return attachment, result.Error
}
