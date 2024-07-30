package main

import (
	_ "github.com/yelnar0112/project-management/docs"

	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/yelnar0112/project-management/internal/config"
	"github.com/yelnar0112/project-management/internal/handler"
)

// @title Project Management API
// @version 1.0
// @description This is a sample server for a project management system.
// @termsOfService http://swagger.io/terms/

// @host localhost:8080
// @BasePath /
func main() {
	config.LoadConfig()

	config.ConnectDB()

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userGroup := router.Group("/user")
	{
		userGroup.GET("/", handler.GetUsers)
		userGroup.POST("/", handler.CreateUser)
		userGroup.GET("/:id", handler.GetUser)
		userGroup.PUT("/:id", handler.UpdateUser)
		userGroup.DELETE("/:id", handler.DeleteUser)
	}

	taskGroup := router.Group("/tasks")
	{
		taskGroup.GET("/", handler.GetTasks)
		taskGroup.POST("/", handler.CreateTask)
		taskGroup.GET("/:id", handler.GetTask)
		taskGroup.PUT("/:id", handler.UpdateTask)
		taskGroup.DELETE("/:id", handler.DeleteTask)
	}

	projectGroup := router.Group("/projects")
	{
		projectGroup.GET("/", handler.GetProjects)
		projectGroup.POST("/", handler.CreateProject)
		projectGroup.GET("/:id", handler.GetProject)
		projectGroup.PUT("/:id", handler.UpdateProject)
		projectGroup.DELETE("/:id", handler.DeleteProject)
	}

	log.Println("Server is running on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}
