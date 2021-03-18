package user

import (
	"time"

	"github.com/google/uuid"
)

// User holds the data to identify a customer.
type User struct {
	Firstname string        `json:"first_name" validate:"required"`
	Lastname  string        `json:"last_name"`
	Email     string        `json:"email" validate:"required,email"`
	Password  string        `json:"password" validate:"required, min=8"`
	Address   Address       `json:"address" validate:"required"`
	UUID      uuid.UUID     `json:"uuid" validate:"required"`
	CreatedAt time.Duration `json:"created_at"`
}

// Address is formed up of City, State and Zipcode.
type Address struct {
	Street  string `json:"street" validate:"required"`
	City    string `json:"city" valdiate:"required"`
	State   string `json:"state" validate:"required"`
	Zipcode int32  `json:"zip_code" validate:"required, len=6"`
}
