package models

type Payment struct {
	// using pointer to int to correctly handle required tag
	Transaction  string `json:"transaction" validate:"required"`
	RequestID    string `json:"request_id"` // optional
	Currency     string `json:"currency" validate:"required,alpha"`
	Provider     string `json:"provider" validate:"required"`
	Amount       *int   `json:"amount" validate:"required,gte=0"`
	PaymentDT    *int64 `json:"payment_dt" validate:"required,gt=0"`
	Bank         string `json:"bank" validate:"required"`
	DeliveryCost *int   `json:"delivery_cost" validate:"required,gte=0"`
	GoodsTotal   *int   `json:"goods_total" validate:"required,gte=0"`
	CustomFee    *int   `json:"custom_fee" validate:"required,gte=0"`
}
