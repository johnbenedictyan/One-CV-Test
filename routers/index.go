package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes add all routing list here automatically get main router
func RegisterRoutes(route *gin.Engine) {
	route.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})
	route.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"live": "ok"}) })

	// Add Api route
	ApiRoutes(route)

	// Add Seed route
	SeedRoutes(route)

	// Add Teacher route
	TeacherRoutes(route)

	// Add Student route
	StudentRoutes(route)
}
