package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johnbenedictyan/One-CV-Test/infra/database"
	"github.com/johnbenedictyan/One-CV-Test/infra/logger"
	"github.com/johnbenedictyan/One-CV-Test/models"
)

type StudentController struct{}

func (ctrl *StudentController) CreateStudent(ctx *gin.Context) {
	student := new(models.Student)

	err := ctx.ShouldBindJSON(&student)
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = database.DB.Create(&student).Error
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &student)
}

func (ctrl *StudentController) GetStudentData(ctx *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	ctx.JSON(http.StatusOK, gin.H{"data": students})

}

func (ctrl *StudentController) GetStudentDataByID(ctx *gin.Context) {
	var student models.Student
	id := ctx.Param("id")
	database.DB.First(&student, id)
	ctx.JSON(http.StatusOK, gin.H{"data": student})
}

func (ctrl *StudentController) UpdateStudentData(ctx *gin.Context) {
	var student models.Student
	id := ctx.Param("id")
	database.DB.First(&student, id)
	ctx.ShouldBindJSON(&student)
	database.DB.Save(&student)
	ctx.JSON(http.StatusOK, gin.H{"data": student})
}

func (ctrl *StudentController) DeleteStudentData(ctx *gin.Context) {
	var student models.Student
	id := ctx.Param("id")
	database.DB.Delete(&student, id)
	ctx.JSON(http.StatusOK, gin.H{"data": true})
}

func (ctrl *StudentController) GetStudentDataByTeacherID(ctx *gin.Context) {
	var students []models.Student
	id := ctx.Param("id")
	database.DB.Where("teacher_id = ?", id).Find(&students)
	ctx.JSON(http.StatusOK, gin.H{"data": students})
}
