package main

import (
	"fmt"
	"github.com/darth-raijin/psql-integration-test-poc/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&models.User{},
		&models.ApiKey{},
		&models.Jobs{}); err != nil {
		return fmt.Errorf("failed to migrate user: %v", err)
	}

	return nil
}
