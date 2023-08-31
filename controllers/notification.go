package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johnbenedictyan/One-CV-Test/infra/database"
	"github.com/johnbenedictyan/One-CV-Test/infra/logger"
	"github.com/johnbenedictyan/One-CV-Test/models"
)

type NotificationController struct{}

func (ctrl *NotificationController) CreateNotification(ctx *gin.Context) {
	notification := new(models.Notification)

	err := ctx.ShouldBindJSON(&notification)
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = database.DB.Create(&notification).Error
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &notification)
}

func (ctrl *NotificationController) GetNotificationData(ctx *gin.Context) {
	var notifications []models.Notification
	database.DB.Find(&notifications)
	ctx.JSON(http.StatusOK, gin.H{"data": notifications})

}

func (ctrl *NotificationController) GetNotificationDataByID(ctx *gin.Context) {
	var notification models.Notification
	id := ctx.Param("id")
	database.DB.First(&notification, id)
	ctx.JSON(http.StatusOK, gin.H{"data": notification})
}

func (ctrl *NotificationController) UpdateNotificationData(ctx *gin.Context) {
	var notification models.Notification
	id := ctx.Param("id")
	database.DB.First(&notification, id)
	ctx.ShouldBindJSON(&notification)
	database.DB.Save(&notification)
	ctx.JSON(http.StatusOK, gin.H{"data": notification})
}

func (ctrl *NotificationController) DeleteNotificationData(ctx *gin.Context) {
	var notification models.Notification
	id := ctx.Param("id")
	database.DB.Delete(&notification, id)
	ctx.JSON(http.StatusOK, gin.H{"data": true})
}

func (ctrl *NotificationController) GetNotificationDataByTeacherID(ctx *gin.Context) {
	var notification []models.Notification
	teacher_id := ctx.Param("teacher_id")
	database.DB.Where("teacher_id = ?", teacher_id).Find(&notification)
	ctx.JSON(http.StatusOK, gin.H{"data": notification})
}

func (ctrl *NotificationController) UsersThatCanReceive(ctx *gin.Context) {
	var notification models.Notification
	id := ctx.Param("id")
	database.DB.First(&notification, id)
	var users []models.User
	database.DB.Model(&notification).Association("Users").Find(&users)
	ctx.JSON(http.StatusOK, gin.H{"data": users})
}
