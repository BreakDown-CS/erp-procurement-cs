package dto

import "github.com/google/uuid"

type StaffSuppliersResponse struct {
	SupplierID uuid.UUID `json:"supplier_id"`
}
