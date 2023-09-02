package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/johnbenedictyan/One-CV-Test/controllers"
)

func SeedRoutes(route *gin.Engine) {
	var teacherCtrl controllers.TeacherController
	var studentCtrl controllers.StudentController
	seed := route.Group("/seed")
	seed.POST("/students", studentCtrl.SeedStudents)
	seed.POST("/teachers", teacherCtrl.SeedTeachers)
}
