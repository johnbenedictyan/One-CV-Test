package models

type Student struct {
	Email string `json:"email" gorm:"primaryKey"`
}
