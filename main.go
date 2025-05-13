package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/yourusername/taskmanager-back/database"
	"github.com/yourusername/taskmanager-back/models"
	"github.com/yourusername/taskmanager-back/routes"
)

func main() {
	// загрузка .env
	godotenv.Load()

	// подключение к БД
	database.Connect()

	// авто-миграция моделей
	db := database.DB
	db.AutoMigrate(
		&models.User{},
		&models.Project{},
		&models.Task{},
		// &models.Tag{}, &models.Comment{}, ...
	)

	r := gin.Default()
	// включим CORS по необходимости
	// r.Use(cors.Default())

	// инициализируем маршруты
	routes.Setup(r)

	r.Run(":8080")
}
