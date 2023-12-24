package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"github.com/woozie-10/students-rest-api/handlers"
)

var Router *gin.Engine = gin.Default()

func InitRouter() {
	RegisterStudentRoutes()
}

func RegisterStudentRoutes() {
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	studentGroup := Router.Group("/students")

	studentGroup.GET("", handlers.GetStudentsHandler)
	studentGroup.GET("/:username", handlers.GetStudentByUsernameHandler)
	studentGroup.GET("/group/:group", handlers.GetStudentsByGroupHandler)
	studentGroup.GET("/course/:course", handlers.GetStudentsByCourseHandler)

	studentGroup.POST("", handlers.AddStudentHandler)
	studentGroup.PUT("/:username", handlers.UpdateStudentHandler)

	studentGroup.DELETE("/:username", handlers.DeleteStudentByUsernameHandler)
}
