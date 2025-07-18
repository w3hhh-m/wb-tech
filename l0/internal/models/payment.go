package models

// Payment contains payment info.
// @Description Payment details for the order.
type Payment struct {
	// Transaction ID
	Transaction string `json:"transaction" validate:"required"`
	// Request ID
	RequestID string `json:"request_id"`
	// Payment currency
	Currency string `json:"currency" validate:"required,alpha"`
	// Payment provider
	Provider string `json:"provider" validate:"required"`
	// Payment amount
	Amount *int `json:"amount" validate:"required,gte=0"`
	// Payment timestamp
	PaymentDT *int64 `json:"payment_dt" validate:"required,gt=0"`
	// Bank name
	Bank string `json:"bank" validate:"required"`
	// Delivery cost
	DeliveryCost *int `json:"delivery_cost" validate:"required,gte=0"`
	// Total goods cost
	GoodsTotal *int `json:"goods_total" validate:"required,gte=0"`
	// Custom fee
	CustomFee *int `json:"custom_fee" validate:"required,gte=0"`
}
