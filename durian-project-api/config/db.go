package config

import (
	"log"
	"os"
	"time"

	"github.com/bryantang1107/Jom-Fresh/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // Use the standard Go logger
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color output
		},
	)
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/postgres"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	// Clear existing tables if needed (Be cautious with this in production)
	// db.Migrator().DropTable(&models.OrderItem{}, &models.Order{}, &models.User{}, &models.Durian{})

	// Perform migrations in the recommended order
	err = db.AutoMigrate(&models.Durian{})
	if err != nil {
		log.Fatalf("failed to migrate Durian: %v", err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("failed to migrate User: %v", err)
	}

	err = db.AutoMigrate(&models.Order{})
	if err != nil {
		log.Fatalf("failed to migrate Order: %v", err)
	}

	err = db.AutoMigrate(&models.OrderItem{})
	if err != nil {
		log.Fatalf("failed to migrate OrderItem: %v", err)
	}
	DB = db
}
