package database

import (
	"backend/pkg/logger"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type DB struct {
	Database *gorm.DB
}

func DBConnection(dsn string) (*DB, error) {

	// Determine logging level and initialize logging
	logMode := viper.GetBool("DB_LOGMODE")
	loglevel := gormLogger.Silent
	if logMode {
		loglevel = gormLogger.Info
	}

	// Build Connection to database
	pgConn := postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	})

	db, err := gorm.Open(pgConn, &gorm.Config{Logger: gormLogger.Default.LogMode(loglevel)})

	if err != nil {
		logger.Fatal("database refused %v", err)
	}
	logger.Log("Connected to database")

	return &DB{Database: db}, nil
}

func WatchDBConnection(db *DB) error {

	postgresDB, err := db.Database.DB()
	timeoutCount := 0
	if err != nil {
		return err
	}
	for {
		error := postgresDB.Ping()
		if error != nil {
			timeoutCount += 1
			if timeoutCount == 5 {
				logger.Fatal("Connection to database failed after trying to reconnect 5 times. Stopping application ")
			}
			logger.Error("Connection to database lost. Try to reconnect %v", timeoutCount)
			time.Sleep(time.Second * 5)
			continue
		}
		timeoutCount = 0
		time.Sleep(time.Second * 5)
	}
}
