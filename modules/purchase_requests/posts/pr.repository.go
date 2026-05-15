package posts

import (
	"context"

	"github.com/BreakDown-CS/erp-procurement-cs/modules/purchase_requests/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type PurchaseRequestRepository interface {
	GeneratePRNO(ctx context.Context, tx pgx.Tx) (string, error)
	InsertPurchaseRequest(ctx context.Context, tx pgx.Tx, supplier model.PurchaseRequests) (uuid.UUID, error)
	InsertPurchaseRequestDetail(ctx context.Context, tx pgx.Tx, supplier model.PurchaseRequestDetails) (uuid.UUID, error)
	InsertPurchaseRequestApproved(ctx context.Context, tx pgx.Tx, supplier model.PurchaseRequestApproved) (uuid.UUID, error)
}
