package models

type Teacher struct {
	ID                 uint      `gorm:"primary_key" json:"-"`
	Email              string    `gorm:"unique;not null" json:"email"`
	RegisteredStudents []Student `gorm:"many2many:teacher_students;" json:"students"`
}

// TableName Database Table Name of this model
func (e *Teacher) TableName() string {
	return "teachers"
}
