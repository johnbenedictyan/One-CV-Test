package models

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Email        string     `gorm:"unique;not null;serializer:json" json:"email"`
	IsSuspended  bool       `gorm:"default:false" json:"suspended"`
	RegisteredTo []*Teacher `gorm:"many2many:teacher_students;" json:"teachers"`
}

// TableName Database Table Name of this model
func (e *Student) TableName() string {
	return "students"
}
