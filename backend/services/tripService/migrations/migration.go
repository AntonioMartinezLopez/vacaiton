package migrations

import (
	"backend/pkg/database"
	"backend/services/tripService/models"
)

func Migrate(db *database.DB) {
	var migrationModels = []interface{}{&models.Trip{}, &models.TripQuery{}, &models.TripStop{}, &models.StopHighlight{}}
	err := db.Database.AutoMigrate(migrationModels...)
	if err != nil {
		return
	}
}
