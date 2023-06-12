package entity

import "gorm.io/gorm"

type Module struct {
	*gorm.Model
	ModuleName   string       `json:"module_name" form:"module_name"`
	SectionId    string       `json:"section_id" form:"section_id"`
	Section      Section      `json:"section" gorm:"foreignKey:SectionId"`
	AttachmentId *string      `json:"attachment_id" form:"attachment_id"`
	Attachment   Attachment   `json:"attachment,omitempty" gorm:"foreignKey:AttachmentId"`
	Tasks        []Task       `json:"tasks" gorm:"foreignKey:ModuleId"`
	Submission   []Submission `json:"submission" gorm:"foreignKey:ModuleId"`
}
