package storage

import (
	"fmt"
	"github.com/stranik28/HackLoaylityProgramm/internal/helper"
	"github.com/stranik28/HackLoaylityProgramm/internal/logger"
	"github.com/stranik28/HackLoaylityProgramm/internal/models"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the database connection.
var DB *gorm.DB

// SetupDatabase migrates and sets up the database.
func SetupDatabase() {
	user := helper.GetEnv("POSTGRES_USER", "golang")
	password := helper.GetEnv("POSTGRES_PASSWORD", "golang")
	host := helper.GetEnv("POSTGRES_HOST", "localhost")
	port := helper.GetEnv("POSTGRES_PORT", "5432")
	dbName := helper.GetEnv("POSTGRES_DB", "go_test")

	// Assemble the connection string.
	psqlSetup := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbName, password)

	logger.Log.Info("Connecting to PostgresSQL...", zap.String("psqlSetup", psqlSetup))

	db, err := gorm.Open(postgres.Open(psqlSetup))

	if err != nil {
		logger.Log.Fatal("Failed to connect to database")
		panic("Could not open database connection")
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		logger.Log.Fatal("Database migration failed", zap.Error(err))
		panic(err)
	}

	DB = db
}
