package router

import (
	"github.com/gin-gonic/gin"
	"github.com/woozie-10/students-rest-api/handlers"
)

var Router *gin.Engine = gin.Default()

func InitRouter() {
	RegisterStudentRoutes()
}

func RegisterStudentRoutes() {
	studentGroup := Router.Group("/students")

	studentGroup.GET("", handlers.GetStudentsHandler)
	studentGroup.GET("/:username", handlers.GetStudentbyUsernameHandler)
	studentGroup.GET("/group/:group", handlers.GetStudentsByGroupHandler)
	studentGroup.GET("/course/:course", handlers.GetStudentsByCourseHandler)

	studentGroup.POST("", handlers.AddStudentHandler)
	studentGroup.PUT("/:username", handlers.UpdateStudentHandler)

	studentGroup.DELETE("/:username", handlers.DeleteStudentByUsernameHandler)
}
