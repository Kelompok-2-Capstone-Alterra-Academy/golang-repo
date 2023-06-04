package entity

import "gorm.io/gorm"

type Folder struct {
	*gorm.Model

	FolderName string `json:"folder_name" form:"folder_name" validate:"required"`
	MentorId   int    `json:"mentor_id"`
	Mentor     User   `json:"mentor" gorm:"foreignKey:MentorId"`
}
