package database

import (
	"errors"
	"fmt"

	"weekly-newsletter/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(model.User{})
}

func ConnectPostgres(host string, port string, username string, pass string, dbname string) (*gorm.DB, error) {
	if host == "" && port == "" && dbname == "" {
		return nil, errors.New("cannot established the connection")
	}
	connectionStr := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", host, username, pass, dbname, port)

	gormConfig := &gorm.Config{
		DryRun: false,
		Logger: logger.Default.LogMode(logger.Info),
	}
	db, err := gorm.Open(postgres.Open(connectionStr), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("connect postgres: %w", err)
	}

	return db, nil
}

func DisconnectPostgres(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("disconnect postgres: %w", err)
	}

	return sqlDB.Close()
}
