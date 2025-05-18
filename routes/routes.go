package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pavhol/taskmanager-back/controllers"
	"github.com/pavhol/taskmanager-back/middleware"
)

func Setup(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api")

	// Auth
	auth := api.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
		auth.GET("/profile", middleware.AuthMiddleware(), controllers.Profile)

	}

	// Защищённые маршруты
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{

		// Пользователи
		users := protected.Group("/users")
		{
			users.GET("", controllers.GetUsers)
			users.GET("/:id", controllers.GetUser)
			users.PUT("/:id", controllers.UpdateUser)
			users.DELETE("/:id", controllers.DeleteUser)
		}

		// Проекты
		projects := protected.Group("/projects")
		{
			projects.POST("", controllers.CreateProject)
			projects.GET("", controllers.GetProjects)
			projects.GET("/:id", controllers.GetProject)
			projects.PUT("/:id", controllers.UpdateProject)
			projects.DELETE("/:id", controllers.DeleteProject)
		}

		// Задачи
		tasks := protected.Group("/tasks")
		{
			tasks.POST("", controllers.CreateTask)
			tasks.GET("", controllers.GetTasks)
			tasks.GET("/:id", controllers.GetTask)
			tasks.PUT("/:id", controllers.UpdateTask)
			tasks.DELETE("/:id", controllers.DeleteTask)
			tasks.GET("/me", controllers.GetTasks)
		}
	}
}
