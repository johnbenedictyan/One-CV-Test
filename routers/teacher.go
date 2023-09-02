package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/johnbenedictyan/One-CV-Test/controllers"
)

func TeacherRoutes(route *gin.Engine) {
	var teacherCtrl controllers.TeacherController
	teacher := route.Group("/teacher")
	teacher.POST("/create", teacherCtrl.CreateTeacher)
	teacher.GET("/get", teacherCtrl.GetTeacherData)
	teacher.GET("/get/:id", teacherCtrl.GetTeacherDataByID)
	teacher.GET("/get/email/:email", teacherCtrl.GetTeacherDataByEmail)
	teacher.PUT("/update/:id", teacherCtrl.UpdateTeacherData)
	teacher.DELETE("/delete/:id", teacherCtrl.DeleteTeacherData)
}
