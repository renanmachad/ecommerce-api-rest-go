package di

import (
	"github.com/renanmachad/ecommerce-rest/pkg/config/database"
	"gorm.io/gorm"
)

func NewConfig() (*gorm.DB, error) {
	return database.Connect()
}
