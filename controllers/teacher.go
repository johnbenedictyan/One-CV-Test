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

func (ctrl *TeacherController) RegisterStudents(ctx *gin.Context) {
	var teacher models.Teacher
	id := ctx.Param("id")
	database.DB.First(&teacher, id)
	var student []models.Student
	ctx.ShouldBindJSON(&student)
	database.DB.Model(&teacher).Association("Students").Append(&student)
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

func (ctrl *TeacherController) SuspendStudent(ctx *gin.Context) {
	var student models.Student
	id := ctx.Param("id")
	database.DB.First(&student, id)
	student.IsSuspended = true
	database.DB.Save(&student)
	ctx.JSON(http.StatusOK, gin.H{"data": student})
}
