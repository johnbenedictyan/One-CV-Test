package models

type Teacher struct {
	ID                 uint           `gorm:"primary_key" json:"id"`
	Email              string         `gorm:"unique;not null" json:"email"`
	RegisteredStudents []Student      `gorm:"many2many:teacher_students;"`
	Notifications      []Notification `json:"notifications"`
}

// TableName Database Table Name of this model
func (e *Teacher) TableName() string {
	return "teachers"
}
