package models

type Student struct {
	ID           uint      `gorm:"primary_key" json:"-"`
	Email        string    `gorm:"unique;not null" json:"email"`
	Suspended    bool      `gorm:"default:false" json:"-"`
	RegisteredTo []Teacher `gorm:"many2many:teacher_students;" json:"-"`
}

// TableName Database Table Name of this model
func (e *Student) TableName() string {
	return "students"
}
