package entity

import "gorm.io/gorm"

type Attachment struct {
	*gorm.Model

	AttachmentName   string  `json:"attachment_name" form:"attachment_name" `
	Description      string  `json:"description" form:"description" `
	Type             string  `json:"type" form:"type" `
	AttachmentSource string  `json:"attachment_source" form:"attachment_source" `
	Status           string  `json:"status" gorm:"default:draft"  form:"status"`
	FolderId         *string `json:"folder_id" form:"folder_id"`
	Folder           Folder  `json:"folder,omitempty" gorm:"foreignKey:FolderId"`
}
