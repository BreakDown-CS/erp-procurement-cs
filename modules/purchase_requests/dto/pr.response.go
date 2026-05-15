package dto

import "github.com/google/uuid"

type PurchaseRequestResponse struct {
	PRId         uuid.UUID   `json:"pr_id"`
	PRDetailId   []uuid.UUID `json:"pr_detail_id"`
	PRApprovedId []uuid.UUID `json:"pr_approved_id"`
}
