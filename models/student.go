package models

type Student struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	Email        string    `gorm:"unique;not null" json:"email"`
	IsSuspended  bool      `gorm:"default:false" json:"suspended"`
	RegisteredTo []Teacher `gorm:"many2many:teacher_students;"`
}

// TableName Database Table Name of this model
func (e *Student) TableName() string {
	return "students"
}
