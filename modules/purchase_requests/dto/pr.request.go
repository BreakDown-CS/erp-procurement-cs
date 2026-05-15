package dto

import "github.com/google/uuid"

type CreatPurchaseRequest struct {
	RequestDate  string    `json:"request_date" validate:"required"`
	StaffId      uuid.UUID `json:"staff_id" validate:"required"`
	DepartmentId uuid.UUID `json:"department_id" validate:"required"`
	Remark       string    `json:"remark" validate:"required"`
	Items        []CreatePurchaseRequestItem
}

type CreatePurchaseRequestItem struct {
	ProductID uuid.UUID `json:"product_id" validate:"required"`
	Qty       int       `json:"qty" validate:"required"`
	UnitPrice int       `json:"unit_price" validate:"required"`
}
