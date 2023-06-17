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

func (repo AttachmentRepository) GetAttachment(id int) (entity.Attachment, error) {
	var Attachments entity.Attachment
	result := repo.DB.First(&Attachments, id)
	return Attachments, result.Error
}

func (repo AttachmentRepository) CreateAttachment(Attachment entity.Attachment) error {
	result := repo.DB.Create(&Attachment)
	return result.Error
}

func (repo AttachmentRepository) UpdateAttachment(id int, Attachment entity.Attachment) error {
	result := repo.DB.Model(&Attachment).Where("id = ?", id).Updates(&Attachment)
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

	// Query the attachments with type "materi"
	if err := repo.DB.Where("type = ?", "video").Find(&attachments).Error; err != nil {
		// Handle the error
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
