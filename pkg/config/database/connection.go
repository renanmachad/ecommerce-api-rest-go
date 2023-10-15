package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// TODO:

func Connect() (*gorm.DB, error) {
	dsn := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable&search_path=goland"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "goland.",
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	return db, err
}
