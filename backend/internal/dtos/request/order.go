package request

import "github.com/google/uuid"

type OrderRequest struct {
	AddressID     uuid.UUID `json:"address_id"`
	CartID        uuid.UUID `json:"cart_id"`
	PaymentMethod string    `json:"payment_method"`
}

type UpdateOrderStatusRequest struct {
	OrderID uuid.UUID `json:"order_id"`
	Status  string    `json:"status"`
}
