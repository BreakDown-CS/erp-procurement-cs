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
	var suppliersId uuid.UUID

	querySupplierInsert := `
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
		querySupplierInsert,
		supplier.PRNo,
		supplier.StaffRequestId,
		supplier.DepartmentID,
		supplier.Remark,
		supplier.CreatedBy,
	).Scan(&suppliersId)

	if err != nil {
		return uuid.Nil, err
	}

	return suppliersId, nil
}
