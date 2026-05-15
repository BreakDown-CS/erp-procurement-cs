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

func (r *repository) GeneratePRNO(ctx context.Context, tx pgx.Tx, departuentsId uuid.UUID) (string, error) {
	var prNo string

	queryGeneratePRNO := `
		SELECT
			'PR' || to_char(NOW(), 'yyyymm') || '/' || COALESCE (
				REPLACE(
				to_char(
					(
					substr(MAX(pr.pr_no), LENGTH(MAX(pr.pr_no)) - 3, 4) :: INT + 1
					),
					'FM0000'
				),
				' ',
				''
				),
				'0001'
			) AS reqno
		FROM
			erp.departments dpm
			LEFT JOIN erp.purchase_requests pr ON pr.department_id = dpm.id AND to_char(pr.request_date, 'mmyyyy') = to_char(NOW(), 'mmyyyy')
		WHERE
			dpm.id = $1
	`

	err := tx.QueryRow(ctx, queryGeneratePRNO, departuentsId).Scan(&prNo)

	if err != nil {
		return "", err
	}

	return prNo, nil
}

func (r *repository) InsertPurchaseRequest(ctx context.Context, tx pgx.Tx, pr model.PurchaseRequests) (uuid.UUID, error) {
	var purchaseRequestId uuid.UUID

	queryPurchaseRequestInsert := `
		INSERT INTO erp.purchase_requests (
			pr_no,
			request_date,
			staff_request_id,
			department_id,
			remark,
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
		queryPurchaseRequestInsert,
		pr.PRNo,
		pr.RequestDate,
		pr.StaffRequestId,
		pr.DepartmentID,
		pr.Remark,
		pr.CreatedBy,
	).Scan(&purchaseRequestId)

	if err != nil {
		return uuid.Nil, err
	}

	return purchaseRequestId, nil
}

func (r *repository) InsertPurchaseRequestDetail(ctx context.Context, tx pgx.Tx, pr model.PurchaseRequestDetails) (uuid.UUID, error) {
	var purchaseRequestDetailId uuid.UUID

	queryPurchaseRequestDetailInsert := `
		INSERT INTO erp.purchase_request_details (
			purchase_request_id,
			product_id,
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
		pr.PurchaseRequestID,
		pr.ProductId,
		pr.Qty,
		pr.UnitPrice,
		pr.TotalPrice,
		pr.CreatedBy,
	).Scan(&purchaseRequestDetailId)

	if err != nil {
		return uuid.Nil, err
	}

	return purchaseRequestDetailId, nil
}

func (r *repository) InsertPurchaseRequestApproved(ctx context.Context, tx pgx.Tx, pr model.PurchaseRequestApproved) (uuid.UUID, error) {
	var purchaseRequestApprovedId uuid.UUID

	queryPurchaseRequestApprovedInsert := `
		INSERT INTO erp.purchase_request_approved (
			purchase_request_detail_id,
			updated_by
		) VALUES (
			$1,
			$2
		) RETURNING id
	`

	err := tx.QueryRow(ctx,
		queryPurchaseRequestApprovedInsert,
		pr.PurchaseRequestDetailID,
		pr.UpdatedBy,
	).Scan(&purchaseRequestApprovedId)

	if err != nil {
		return uuid.Nil, err
	}

	return purchaseRequestApprovedId, nil
}
