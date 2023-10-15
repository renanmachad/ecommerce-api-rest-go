package entities

import "time"

type Order struct {
	ID          int64     `json:"id"`
	CustomerID  int       `json:"customer_id"`
	TotalAmount int       `json:"total_amount"`
	OrderStatus string    `json:"order_status"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewOrder(id int64, customerId int, totalAmount int, orderStatus string, createdAt time.Time) *Order {
	return &Order{
		ID:          id,
		CustomerID:  customerId,
		TotalAmount: totalAmount,
		OrderStatus: orderStatus,
		CreatedAt:   createdAt,
	}
}
