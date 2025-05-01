package connection

import (
	"fmt"

	"github.com/7lawa9111/Project-Based-Mentorship-Go-App/models"

	// "github.com/7lawa9111/Project-Based-Mentorship-Go-App/issues/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabaseConnection() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, falling back to system environment variables")
	}
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL_MODE"),
	)

	config := &gorm.Config{}
	if os.Getenv("ENV") == "development" {
		config.Logger = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Auto-migrate schemas
	if err := db.AutoMigrate(&models.Author{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}
	if err := db.AutoMigrate(&models.Document{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}
	log.Println("Database connection established successfully")
	return db, nil
}
