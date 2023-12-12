package migrations

import (
	"backend/pkg/database"
)

func Migrate(db *database.DB) {
	var migrationModels = []interface{}{}
	err := db.Database.AutoMigrate(migrationModels...)
	if err != nil {
		return
	}
}
