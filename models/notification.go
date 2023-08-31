package models

import (
	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	TeacherID uint   `gorm:"not null" json:"teacherId"`
	Text      string `gorm:"not null" json:"text"`
}

// TableName Database Table Name of this model
func (e *Notification) TableName() string {
	return "notifications"
}
