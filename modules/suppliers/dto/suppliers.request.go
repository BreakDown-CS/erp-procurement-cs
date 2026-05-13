package dto

type CreateSuppliersRequest struct {
	SupplierCode string `json:"supplier_code" validate:"required"`
	SupplierName string `json:"supplier_name" validate:"required"`
	TaxID        string `json:"tax_id" validate:"required"`
	Email        string `json:"email" validate:"required"`
	Phone        string `json:"phone" validate:"required"`
	Address      string `json:"address" validate:"required"`
	StaffID      string `json:"staff_id" validate:"required"`
}
