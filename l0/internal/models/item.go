package models

type Item struct {
	// using pointer to int to correctly handle required tag
	ChrtID      *int64 `json:"chrt_id" validate:"required,gt=0"`
	TrackNumber string `json:"track_number" validate:"required"`
	Price       *int   `json:"price" validate:"required,gte=0"`
	RID         string `json:"rid" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Sale        *int   `json:"sale" validate:"required,gte=0,lte=100"`
	Size        string `json:"size" validate:"required"`
	TotalPrice  *int   `json:"total_price" validate:"required,gte=0"`
	NmID        *int64 `json:"nm_id" validate:"required"`
	Brand       string `json:"brand" validate:"required"`
	Status      *int   `json:"status" validate:"required"`
}
