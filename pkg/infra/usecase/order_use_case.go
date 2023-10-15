package usecase

import (
	"github.com/renanmachad/ecommerce-rest/pkg/infra/entities"
	"github.com/renanmachad/ecommerce-rest/pkg/infra/repositories"
)

type OrderUseCase struct {
	repo *repositories.OrderRepository
}

// constructor
func NewOrderUseCase(repo *repositories.OrderRepository) *OrderUseCase {
	return &OrderUseCase{
		repo: repo,
	}
}

func (u *OrderUseCase) Create(model *entities.Order) error {
	return u.repo.Create(model)
}
