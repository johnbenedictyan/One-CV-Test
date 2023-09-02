package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/johnbenedictyan/One-CV-Test/controllers"
)

func StudentRoutes(route *gin.Engine) {
	var studentCtrl controllers.StudentController
	student := route.Group("/student")
	student.POST("/create", studentCtrl.CreateStudent)
	student.GET("/get", studentCtrl.GetStudentData)
	student.GET("/get/:id", studentCtrl.GetStudentDataByID)
	student.PUT("/update/:id", studentCtrl.UpdateStudentData)
	student.DELETE("/delete/:id", studentCtrl.DeleteStudentData)
}
