package entity

import "gorm.io/gorm"

type Submission struct {
	*gorm.Model
	SubmissionSource string  `json:"submission_source" form:"submission_source"`
	Status           string  `json:"status" form:"status"  gorm:"default:belum di nilai"`
	ModuleId         string  `json:"module_id" form:"module_id"`
	Type             string  `json:"type" form:"type"`
	Module           *Module `json:"module,omitempty" gorm:"foreignKey:ModuleId"`
	StudentId        string  `json:"student_id" form:"student_id"`
	User             *User   `json:"student,omitempty" gorm:"foreignKey:StudentId"`
}
