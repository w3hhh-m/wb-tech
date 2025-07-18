package models

import "time"

// Order represents an order entity.
// @Description Order received from external system. Contains delivery, payment and items info.
type Order struct {
	// Unique order identifier
	OrderUID string `json:"order_uid" validate:"required"`
	// Tracking number
	TrackNumber string `json:"track_number" validate:"required"`
	Entry       string `json:"entry" validate:"required"`
	// Delivery info
	Delivery Delivery `json:"delivery" validate:"required"`
	// Payment info
	Payment Payment `json:"payment" validate:"required"`
	// List of items
	Items []Item `json:"items" validate:"required,min=1"`
	// Locale
	Locale string `json:"locale" validate:"required,alpha"`
	// Internal signature
	InternalSignature string `json:"internal_signature"`
	// Customer ID
	CustomerID string `json:"customer_id" validate:"required"`
	// Delivery service
	DeliveryService string `json:"delivery_service" validate:"required"`
	ShardKey        string `json:"shardkey" validate:"required,numeric"`
	SmID            *int   `json:"sm_id" validate:"required,gte=0"`
	// Order creation date
	DateCreated time.Time `json:"date_created" validate:"required"`
	OofShard    string    `json:"oof_shard" validate:"required,numeric"`
}
