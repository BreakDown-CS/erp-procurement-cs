package repository

import (
	"context"

	"github.com/BreakDown-CS/erp-procurement-cs/modules/purchase_requests/model"
	ports "github.com/BreakDown-CS/erp-procurement-cs/modules/purchase_requests/posts"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) ports.PurchaseRequestRepository {
	return &repository{db: db}
}

func (r *repository) GeneratePRNO(ctx context.Context, tx pgx.Tx) (string, error) {
	return "", nil
}

func (r *repository) InsertPurchaseRequest(ctx context.Context, tx pgx.Tx, supplier model.PurchaseRequests) (uuid.UUID, error) {
	var purchaseRequestId uuid.UUID

	queryPurchaseRequestInsert := `
		INSERT INTO erp.purchase_requests (
			pr_no,
			staff_request_id,
			department_id,
			remark,
			created_by
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5
		) RETURNING id
	`

	err := tx.QueryRow(ctx,
		queryPurchaseRequestInsert,
		supplier.PRNo,
		supplier.StaffRequestId,
		supplier.DepartmentID,
		supplier.Remark,
		supplier.CreatedBy,
	).Scan(&purchaseRequestId)

	if err != nil {
		return uuid.Nil, err
	}

	return purchaseRequestId, nil
}

func (r *repository) InsertPurchaseRequestDetail(ctx context.Context, tx pgx.Tx, supplier model.PurchaseRequestDetails) (uuid.UUID, error) {
	var purchaseRequestDetailId uuid.UUID

	queryPurchaseRequestDetailInsert := `
		INSERT INTO erp.purchase_request_details (
			purchase_request_id,
			prodcut_id,
			qty,
			unit_price,
			total_price,
			created_by
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6
		) RETURNING id
	`

	err := tx.QueryRow(ctx,
		queryPurchaseRequestDetailInsert,
		supplier.PurchaseRequestID,
		supplier.ProductId,
		supplier.Qty,
		supplier.TotalPrice,
		supplier.CreatedBy,
	).Scan(&purchaseRequestDetailId)

	if err != nil {
		return uuid.Nil, err
	}

	return purchaseRequestDetailId, nil
}

func (r *repository) InsertPurchaseRequestApproved(ctx context.Context, tx pgx.Tx, supplier model.PurchaseRequestApproved) (uuid.UUID, error) {
	var purchaseRequestApprovedId uuid.UUID

	queryPurchaseRequestApprovedInsert := `
		INSERT INTO erp.purchase_request_approved (
			purchase_request_detail_id,
			created_by
		) VALUES (
			$1,
			$2
		) RETURNING id
	`

	err := tx.QueryRow(ctx,
		queryPurchaseRequestApprovedInsert,
		supplier.PurchaseRequestDetailID,
		supplier.UpdatedBy,
	).Scan(&purchaseRequestApprovedId)

	if err != nil {
		return uuid.Nil, err
	}

	return purchaseRequestApprovedId, nil
}
