// Package database declares a function open a secure connection to database.
package database

import (
	"github.com/Joe-Degs/abeg_mock_backend/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(config.DBNAME), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
