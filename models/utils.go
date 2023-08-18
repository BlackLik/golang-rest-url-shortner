package models

import (
	"log"
	"os"
	"time"

	"golang.org/x/exp/slog"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"urlshort.ru/m/config"
)

var ERROR_HANDLER string = "models"

var DATABASE *gorm.DB

// init initializes the DATABASE variable by calling the InitDB function with the value of the "DB_NAME" environment variable.
//
// No parameters.
// No return types.
func init() {
	DATABASE = InitDB(config.ConfigAll.DB_NAME)
	slog.Debug(ERROR_HANDLER, DATABASE)
}

// InitDB initializes the database connection and returns a pointer to the gorm.DB object.
//
// DB_NAME: the URL of the database.
// Returns: a pointer to the gorm.DB object.
func InitDB(DB_NAME string) *gorm.DB {
	if DB_NAME == "" {
		slog.Error(ERROR_HANDLER, "DB_NAME is empty")
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)

	db, err := gorm.Open(sqlite.Open(DB_NAME), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		slog.Error(ERROR_HANDLER, err)
	}

	Migrate(db)

	return db
}

// Migrate performs database migration.
//
// db: a pointer to a gorm.DB instance.
//
// There is no return type for this function.
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&URL{})
}
