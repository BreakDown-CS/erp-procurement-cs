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

type PurchaseRequests struct {
	ID             uuid.UUID
	PRNo           string
	RequestDate    time.Time
	StaffRequestId uuid.UUID
	DepartmentID   uuid.UUID
	Status         string
	Remark         string
	ApprovedBy     uuid.UUID
	ApprovedAt     time.Time
	CreatedBy      uuid.UUID
	CreatedAt      time.Time
	UpdatedBy      uuid.UUID
	UpdatedAt      time.Time
}

type PurchaseRequestItems struct {
	ID                uuid.UUID
	PurchaseRequestID uuid.UUID
	ProductId         uuid.UUID
	Qty               float64
	UnitPrice         float64
	TotalPrice        float64
	CreatedBy         uuid.UUID
	CreatedAt         time.Time
	UpdatedBy         uuid.UUID
	UpdatedAt         time.Time
}
