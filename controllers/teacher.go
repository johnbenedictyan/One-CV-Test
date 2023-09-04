package controllers

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/johnbenedictyan/One-CV-Test/infra/database"
	"github.com/johnbenedictyan/One-CV-Test/infra/logger"
	"github.com/johnbenedictyan/One-CV-Test/models"
	"gorm.io/gorm"
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
	database.DB.Preload("RegisteredStudents").First(&teacher, id)
	ctx.JSON(http.StatusOK, gin.H{"data": teacher})
}

func (ctrl *TeacherController) GetTeacherDataByEmail(ctx *gin.Context) {
	var teacher models.Teacher
	email := ctx.Param("email")
	println(email)
	database.DB.Where("email = ?", email).First(&teacher)
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
	teacher.RegisteredStudents = append(teacher.RegisteredStudents, students...)
	// Save teacher
	database.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&teacher)

	// Return updated teacher
	var updatedTeacher models.Teacher
	database.DB.Preload("RegisteredStudents").First(&updatedTeacher, teacher.ID)

	ctx.JSON(http.StatusOK, gin.H{"teacher": updatedTeacher})
}

func (ctrl *TeacherController) CommonStudents(ctx *gin.Context) {
	var teachers []models.Teacher
	var students []models.Student
	teacherEmails := ctx.QueryArray("teacher")
	database.DB.Where("email IN ?", teacherEmails).Find(&teachers)
	database.DB.Model(&teachers[0]).Association("RegisteredStudents").Find(&students)

	for i := 1; i < len(teachers); i++ {
		var students2 []models.Student
		database.DB.Model(&teachers[i]).Association("RegisteredStudents").Find(&students2)
		students = intersection(students, students2)
	}
	if students == nil {
		students = []models.Student{}
	}

	ctx.JSON(http.StatusOK, gin.H{"students": students})
}

func intersection(s1 []models.Student, s2 []models.Student) []models.Student {
	// Get intersection of two student arrays by using hash map
	m := make(map[uint]bool)
	for _, student := range s1 {
		m[student.ID] = true
	}
	var intersection []models.Student
	for _, student := range s2 {
		if m[student.ID] {
			intersection = append(intersection, student)
		}
	}

	return intersection
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
	database.DB.Model(&student).Update("suspended", true)

	ctx.JSON(http.StatusOK, gin.H{"student": student})
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

	// Get students registered to teacher only if they are not suspended
	database.DB.Model(&teacher).Association("RegisteredStudents").Find(&students, "suspended = ?", false)

	// Append students with mentioned students
	mentionedStudents := getMentionedStudents(body.Notification)

	for _, studentEmail := range mentionedStudents {
		var student models.Student
		database.DB.Where("email = ?", studentEmail).First(&student, "suspended = ?", false)
		students = append(students, student)
	}

	ctx.JSON(http.StatusOK, gin.H{"students": students})
}

func getMentionedStudents(s string) []string {
	pattern := regexp.MustCompile(`@[a-z0-9]+@[a-z]+\.[a-z]{2,24}`)
	mentionedStudentsEmails := pattern.FindAllString(s, -1)
	for i, studentEmail := range mentionedStudentsEmails {
		mentionedStudentsEmails[i] = studentEmail[1:]
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
