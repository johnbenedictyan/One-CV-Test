package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johnbenedictyan/One-CV-Test/infra/database"
	"github.com/johnbenedictyan/One-CV-Test/infra/logger"
	"github.com/johnbenedictyan/One-CV-Test/models"
)

type TeacherController struct{}

func (ctrl *TeacherController) CreateTeacher(ctx *gin.Context) {
	teacher := new(models.Teacher)

	err := ctx.ShouldBindJSON(&teacher)
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = database.DB.Create(&teacher).Error
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &teacher)
}

func (ctrl *TeacherController) GetTeacherData(ctx *gin.Context) {
	var teachers []models.Teacher
	database.DB.Find(&teachers)
	ctx.JSON(http.StatusOK, gin.H{"data": teachers})

}

func (ctrl *TeacherController) GetTeacherDataByID(ctx *gin.Context) {
	var teacher models.Teacher
	id := ctx.Param("id")
	database.DB.First(&teacher, id)
	ctx.JSON(http.StatusOK, gin.H{"data": teacher})
}

func (ctrl *TeacherController) UpdateTeacherData(ctx *gin.Context) {
	var teacher models.Teacher
	id := ctx.Param("id")
	database.DB.First(&teacher, id)
	ctx.ShouldBindJSON(&teacher)
	database.DB.Save(&teacher)
	ctx.JSON(http.StatusOK, gin.H{"data": teacher})
}

func (ctrl *TeacherController) DeleteTeacherData(ctx *gin.Context) {
	var teacher models.Teacher
	id := ctx.Param("id")
	database.DB.Delete(&teacher, id)
	ctx.JSON(http.StatusOK, gin.H{"data": true})
}

type RegisterStudentsBody struct {
	Teacher  string   `json:"teacher" binding:"required"`
	Students []string `json:"students" binding:"required"`
}

func (ctrl *TeacherController) RegisterStudents(ctx *gin.Context) {
	var teacher models.Teacher
	var body RegisterStudentsBody
	ctx.ShouldBindJSON(&body)
	if err := database.DB.Where("email = ?", body.Teacher).First(&teacher).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Teacher not found"})
		return
	}

	// Get Students from email string array
	var students []models.Student
	database.DB.Where("email IN ?", body.Students).Find(&students)

	// Update teacher's registered students with new students
	database.DB.Model(&teacher).Association("Students").Append(&students)

	ctx.JSON(http.StatusOK, gin.H{"data": teacher})
}

func (ctrl *TeacherController) CommonStudents(ctx *gin.Context) {
	var teacher []models.Teacher
	var students []models.Student
	teacherId := ctx.QueryArray("teacher")
	database.DB.Where("id IN ?", teacherId).Find(&teacher)
	database.DB.Model(&teacher).Association("Students").Find(&students)
	ctx.JSON(http.StatusOK, gin.H{"data": students})
}

type SuspendStudentBody struct {
	Student string `json:"student" binding:"required"`
}

func (ctrl *TeacherController) SuspendStudent(ctx *gin.Context) {
	var student models.Student
	var body SuspendStudentBody
	ctx.ShouldBindJSON(&body)
	if err := database.DB.Where("email = ?", body.Student).First(&student).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Student not found"})
		return
	}

	// Update Student's isSuspended to true
	database.DB.Model(&student).Update("is_suspended", true)

	ctx.JSON(http.StatusOK, gin.H{"data": student})
}

type ListRecipientsBody struct {
	Teacher      string `json:"teacher" binding:"required"`
	Notification string `json:"notification" binding:"required"`
}

func (ctrl *TeacherController) ListRecipients(ctx *gin.Context) {
	var teacher models.Teacher
	var students []models.Student
	var body ListRecipientsBody
	ctx.ShouldBindJSON(&body)

	// Get teacher from email
	if err := database.DB.Where("email = ?", body.Teacher).First(&teacher).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Teacher not found"})
		return
	}

	// Get students registered to teacher
	database.DB.Model(&teacher).Association("Students").Find(&students)

	// Append students with mentioned students
	mentionedStudents := getMentionedStudents(body.Notification)
	for _, studentEmail := range mentionedStudents {
		var student models.Student
		database.DB.Where("email = ?", studentEmail).First(&student)
		students = append(students, student)
	}

}

func getMentionedStudents(s string) []string {
	var mentionedStudentsEmails []string
	// Search for @ in string, append email after @ to mentionedStudentsEmails
	for i := 0; i < len(s); i++ {
		if s[i] == '@' {
			var email string
			for j := i + 1; j < len(s); j++ {
				if s[j] == ' ' {
					break
				}
				email += string(s[j])
			}
			mentionedStudentsEmails = append(mentionedStudentsEmails, email)
		}
	}
	return mentionedStudentsEmails
}

// Seed Teachers given a list of email strings in the request body, creating new teachers if they do not exist
func (ctrl *TeacherController) SeedTeachers(ctx *gin.Context) {
	var body struct {
		Emails []string `json:"emails" binding:"required"`
	}
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var teachers []models.Teacher
	for _, email := range body.Emails {
		var teacher models.Teacher
		if err := database.DB.Where("email = ?", email).First(&teacher).Error; err != nil {
			teacher.Email = email
			database.DB.Create(&teacher)
		}
		teachers = append(teachers, teacher)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": teachers})
}
