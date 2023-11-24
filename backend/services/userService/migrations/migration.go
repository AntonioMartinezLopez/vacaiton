package migrations

import (
	"backend/pkg/database"
	"backend/services/userService/models"
)

func Migrate(db *database.DB) {
	var migrationModels = []interface{}{&models.User{}}
	err := db.Database.AutoMigrate(migrationModels...)
	if err != nil {
		return
	}
}
