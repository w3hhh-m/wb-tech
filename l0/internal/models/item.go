package models

// Item represents a product in the order.
// @Description Product details in the order.
type Item struct {
	ChrtID *int64 `json:"chrt_id" validate:"required,gt=0"`
	// Tracking number
	TrackNumber string `json:"track_number" validate:"required"`
	// Unit price
	Price *int `json:"price" validate:"required,gte=0"`
	// RID
	RID string `json:"rid" validate:"required"`
	// Product name
	Name string `json:"name" validate:"required"`
	// Discount (%)
	Sale *int `json:"sale" validate:"required,gte=0,lte=100"`
	// Size
	Size string `json:"size" validate:"required"`
	// Total price
	TotalPrice *int   `json:"total_price" validate:"required,gte=0"`
	NmID       *int64 `json:"nm_id" validate:"required"`
	// Brand
	Brand string `json:"brand" validate:"required"`
	// Status
	Status *int `json:"status" validate:"required"`
}
