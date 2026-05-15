package model

import (
	"time"

	"github.com/google/uuid"
)

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

type PurchaseRequestDetails struct {
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

type PurchaseRequestApproved struct {
	ID                      uuid.UUID
	PurchaseRequestDetailID uuid.UUID
	ApprovedBy              uuid.UUID
	ApprovedAt              time.Time
	CancelBy                uuid.UUID
	CancelAt                time.Time
	CancelRemark            string
	UpdatedBy               uuid.UUID
	UpdatedAt               time.Time
}
