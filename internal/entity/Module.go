package entity

import (
	"database/sql"
	"encoding/json"

	"gorm.io/gorm"
)

type NullInt struct {
	sql.NullInt64
}

func (ni *NullInt) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		ni.Valid = false
		return nil
	}
	err := json.Unmarshal(data, &ni.Int64)
	ni.Valid = (err == nil)
	return err
}

func (ni NullInt) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int64)
}

type Module struct {
	*gorm.Model
	ModuleName   string       `json:"module_name" form:"module_name"`
	Description  string       `json:"description" form:"description"`
	SectionId    int          `json:"section_id" form:"section_id"`
	Section      Section      `json:"section" gorm:"foreignKey:SectionId"`
	AttachmentId NullInt      `json:"attachment_id" form:"attachment_id"`
	Attachment   Attachment   `json:"attachment,omitempty" gorm:"foreignKey:AttachmentId"`
	Tasks        []Task       `json:"tasks" gorm:"foreignKey:ModuleId"`
	Submission   []Submission `json:"submission" gorm:"foreignKey:ModuleId"`
}
