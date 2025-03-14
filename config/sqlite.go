package config

import (
	"os"

	"github.com/alexandrejuniorc/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	// Instance logger
	logger := GetLogger("sqlite")

	DB_PATH := "./db/main.db"

	// Check if the database file exists
	_, err := os.Stat(DB_PATH)
	if os.IsNotExist(err) {
		logger.Infof("database file does not exist. Creating a new one.")

		// Create the database file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("sqlite folder creation error: %v", err)
			return nil, err
		}

		file, err := os.Create(DB_PATH)
		if err != nil {
			logger.Errorf("sqlite file creation error: %v", err)
			return nil, err
		}

		file.Close()
	}

	// Create DB and connect
	db, err := gorm.Open(sqlite.Open(DB_PATH), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	// Migrate the Schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}

	// Return the DB
	return db, nil
}
