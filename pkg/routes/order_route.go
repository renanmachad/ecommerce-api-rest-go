package routes

import (
	"github.com/gorilla/mux"
	"github.com/renanmachad/ecommerce-rest/pkg/app/rest"
	"github.com/renanmachad/ecommerce-rest/pkg/infra/repositories"
	"github.com/renanmachad/ecommerce-rest/pkg/infra/usecase"
	"gorm.io/gorm"
)

func OrderRoute(router *mux.Router, db *gorm.DB) {

	repository := repositories.NewOrderRepository(db)
	usecaseOrder := usecase.NewOrderUseCase(repository)
	controller := rest.NewOrderController(usecaseOrder)
	// only post
	router.HandleFunc("/api/v1/order", controller.CreateOrder).Methods("POST")
}
