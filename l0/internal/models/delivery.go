package models

// Delivery contains delivery info.
// @Description Delivery details for the order.
type Delivery struct {
	// Recipient name
	Name string `json:"name" validate:"required"`
	// Recipient phone
	Phone string `json:"phone" validate:"required,e164|startswith=+"`
	// Postal code
	Zip string `json:"zip" validate:"required,numeric"`
	// City
	City string `json:"city" validate:"required"`
	// Address
	Address string `json:"address" validate:"required"`
	// Region
	Region string `json:"region" validate:"required"`
	// Email
	Email string `json:"email" validate:"required,email"`
}
