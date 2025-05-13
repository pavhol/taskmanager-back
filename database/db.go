package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), // e.g. "localhost"
		os.Getenv("DB_USER"), // e.g. "postgres"
		os.Getenv("DB_PASS"), // ваш пароль
		os.Getenv("DB_NAME"), // "project_tasks_db"
		os.Getenv("DB_PORT"), // e.g. "5432"
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
}
