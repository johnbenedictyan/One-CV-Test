package models

import (
	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	Email              string         `gorm:"unique;not null;serializer:json" json:"email"`
	RegisteredStudents []*Student     `gorm:"many2many:teacher_students;" json:"students"`
	Notifications      []Notification `json:"notifications"`
}

// TableName Database Table Name of this model
func (e *Teacher) TableName() string {
	return "teachers"
}
