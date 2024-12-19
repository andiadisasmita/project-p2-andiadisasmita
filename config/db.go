package config

import (
	"log"
	"os"

	"github.com/a-andiadisasmita/project-p2-andiadisasmita/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitializeDatabase initializes the database connection
func InitializeDatabase() error {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// Auto-migrate all models
	err = db.AutoMigrate(
		&models.User{},
		&models.Boardgame{},
		&models.Stock{},
		&models.RentalHistory{},
		&models.Payment{},
		&models.Review{},
	)
	if err != nil {
		return err
	}

	DB = db
	log.Println("Database connection and migrations successful")
	return nil
}
