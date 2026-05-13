package model

import (
	"time"

	"github.com/google/uuid"
)

type Suppliers struct {
	ID           uuid.UUID
	SupplierCode string
	SupplierName string
	TaxID        string
	Email        string
	Phone        string
	Address      string
	Status       string
	CreatedBy    uuid.UUID
	CreatedAt    time.Time
	UpdatedBy    uuid.UUID
	UpdatedAt    time.Time
}
