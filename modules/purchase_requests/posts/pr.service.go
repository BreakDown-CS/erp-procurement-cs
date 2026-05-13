package posts

import (
	"context"

	"github.com/BreakDown-CS/erp-procurement-cs/modules/purchase_requests/dto"
)

type PurchaseRequestService interface {
	CreatePurchaseRequests(ctx context.Context, supplier dto.CreatPurchaseRequest) (dto.PurchaseRequestResponse, error)
}
