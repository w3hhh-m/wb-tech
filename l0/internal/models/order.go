package models

import "time"

type Order struct {
	// using pointer to int to correctly handle required tag
	OrderUID          string    `json:"order_uid" validate:"required"`
	TrackNumber       string    `json:"track_number" validate:"required"`
	Entry             string    `json:"entry" validate:"required"`
	Delivery          Delivery  `json:"delivery" validate:"required"`
	Payment           Payment   `json:"payment" validate:"required"`
	Items             []Item    `json:"items" validate:"required,min=1"`
	Locale            string    `json:"locale" validate:"required,alpha"`
	InternalSignature string    `json:"internal_signature"` // optional
	CustomerID        string    `json:"customer_id" validate:"required"`
	DeliveryService   string    `json:"delivery_service" validate:"required"`
	ShardKey          string    `json:"shardkey" validate:"required,numeric"`
	SmID              *int      `json:"sm_id" validate:"required,gte=0"`
	DateCreated       time.Time `json:"date_created" validate:"required"`
	OofShard          string    `json:"oof_shard" validate:"required,numeric"`
}
