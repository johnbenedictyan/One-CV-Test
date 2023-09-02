package repositories

import (
	"github.com/johnbenedictyan/One-CV-Test/infra/database"
	"github.com/johnbenedictyan/One-CV-Test/models"
)

type TeacherRepository struct{}

func (repo *TeacherRepository) CreateTeacher(teacher *models.Teacher) error {
	err := database.DB.Create(&teacher).Error
	return err
}

func (repo *TeacherRepository) GetTeacherData() ([]models.Teacher, error) {
	var teachers []models.Teacher
	err := database.DB.Find(&teachers).Error
	return teachers, err
}

func (repo *TeacherRepository) GetTeacherDataByID(id string) (models.Teacher, error) {
	var teacher models.Teacher
	err := database.DB.First(&teacher, id).Error
	return teacher, err
}

func (repo *TeacherRepository) GetTeacherDataByEmail(email string) (models.Teacher, error) {
	var teacher models.Teacher
	err := database.DB.Where("email = ?", email).First(&teacher).Error
	return teacher, err
}

func (repo *TeacherRepository) UpdateTeacherData(teacher *models.Teacher) error {
	err := database.DB.Save(&teacher).Error
	return err
}

func (repo *TeacherRepository) DeleteTeacherData(teacher *models.Teacher) error {
	err := database.DB.Delete(&teacher).Error
	return err
}
