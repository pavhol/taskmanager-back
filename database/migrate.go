package database

import (
	"github.com/pavhol/taskmanager-back/models"
)

func Migrate() {
	db := DB
	db.AutoMigrate(
		&models.User{},
		&models.Project{},
		&models.Task{},
	)
}
