package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/pavhol/taskmanager-back/database"
	"github.com/pavhol/taskmanager-back/models"
	"github.com/pavhol/taskmanager-back/routes"
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

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// инициализируем маршруты
	routes.Setup(r)

	r.Run(":8080")
}
