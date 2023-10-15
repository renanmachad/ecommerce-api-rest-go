package request

type OrderRequest struct {
	CustomerId  int64  `json:"customer_id"`
	TotalAmount int64  `json:"total_amount"`
	OrderStatus string `json:"order_status"`
}

func newOrderRequest(customerId int64, totalAmount int64, orderStatus string) *OrderRequest {
	return &OrderRequest{
		CustomerId:  customerId,
		OrderStatus: orderStatus,
		TotalAmount: totalAmount,
	}
}
