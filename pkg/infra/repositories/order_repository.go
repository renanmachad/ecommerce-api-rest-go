package repositories

import (
	"github.com/renanmachad/ecommerce-rest/pkg/infra/entities"
	"github.com/renanmachad/ecommerce-rest/pkg/utils"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (repository *OrderRepository) Create(order *entities.Order) error {

	result := repository.db.Create(&order)
	utils.ErrorHandler(result.Error)

	return nil
}
