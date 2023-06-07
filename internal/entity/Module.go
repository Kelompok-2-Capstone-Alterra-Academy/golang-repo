package entity

import "gorm.io/gorm"

type Module struct {
	*gorm.Model

	SectionId    string     `json:"section_id" form:"section_id"`
	Section      Section    `json:"section" gorm:"foreignKey:SectionId"`
	AttachmentId string     `json:"attachment_id" form:"attachment_id"`
	Attachment   Attachment `json:"attachment" gorm:"foreignKey:AttachmentId"`
}
