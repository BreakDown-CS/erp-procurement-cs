package posts

import (
	"context"

	"github.com/BreakDown-CS/erp-procurement-cs/modules/purchase_requests/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type PurchaseRequestRepository interface {
	GeneratePRNO(context.Context, pgx.Tx, uuid.UUID) (string, error)
	InsertPurchaseRequest(context.Context, pgx.Tx, model.PurchaseRequests) (uuid.UUID, error)
	InsertPurchaseRequestDetail(context.Context, pgx.Tx, model.PurchaseRequestDetails) (uuid.UUID, error)
	InsertPurchaseRequestApproved(context.Context, pgx.Tx, model.PurchaseRequestApproved) (uuid.UUID, error)
}
