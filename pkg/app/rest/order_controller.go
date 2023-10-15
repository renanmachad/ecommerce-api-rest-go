package rest

import (
	"encoding/json"
	orderRequest "github.com/renanmachad/ecommerce-rest/pkg/app/dto/request"
	"github.com/renanmachad/ecommerce-rest/pkg/app/dto/response"
	"github.com/renanmachad/ecommerce-rest/pkg/infra/entities"
	"github.com/renanmachad/ecommerce-rest/pkg/infra/usecase"
	RenderResponse "github.com/renanmachad/ecommerce-rest/pkg/utils"
	errorH "github.com/renanmachad/ecommerce-rest/pkg/utils"
	"net/http"
)

type OrderController struct {
	usecase *usecase.OrderUseCase
}

// constructor
func NewOrderController(useCase *usecase.OrderUseCase) *OrderController {
	return &OrderController{
		usecase: useCase,
	}
}
func (c *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var request orderRequest.OrderRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	errorH.ErrorHandler(err)

	orderModel := entities.Order{
		OrderStatus: request.OrderStatus,
		TotalAmount: int(request.TotalAmount),
		CustomerID:  int(request.CustomerId),
	}

	errorUseCase := c.usecase.Create(&orderModel)
	errorH.ErrorHandler(errorUseCase)

	successResponse := response.NewSuccessResponse("Order created successfully", []string{}, "success")
	RenderResponse.RenderSuccess(w, r, successResponse)
}
