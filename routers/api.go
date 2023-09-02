package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/johnbenedictyan/One-CV-Test/controllers"
)

func ApiRoutes(route *gin.Engine) {
	var teacherCtrl controllers.TeacherController
	api := route.Group("/api")
	api.POST("/register", teacherCtrl.RegisterStudents)
	api.GET("/commonstudents", teacherCtrl.CommonStudents)
	api.POST("/suspend", teacherCtrl.SuspendStudent)
	api.POST("/retrievefornotifications", teacherCtrl.ListRecipients)
}
