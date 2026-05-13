package service

import (
	"context"

	"github.com/BreakDown-CS/erp-procurement-cs/internal/helper"
	"github.com/BreakDown-CS/erp-procurement-cs/modules/purchase_requests/dto"
	ports "github.com/BreakDown-CS/erp-procurement-cs/modules/purchase_requests/posts"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type service struct {
	repo   ports.PurchaseRequestRepository
	helper *helper.Uow
	db     *pgxpool.Pool
}

func NewService(repo ports.PurchaseRequestRepository, db *pgxpool.Pool) ports.PurchaseRequestService {
	return &service{
		repo:   repo,
		helper: helper.New(db),
		db:     db,
	}
}

func (s *service) CreatePurchaseRequests(ctx context.Context, supplier dto.CreatPurchaseRequest) (dto.PurchaseRequestResponse, error) {

	response := dto.PurchaseRequestResponse{}

	err := s.helper.WithTx(ctx, func(tx pgx.Tx) error {

		// modelPurchaseRequests := model.PurchaseRequests{
		// 	PRNo: response.PRNo,
		// 	StaffRequestId: response.,
		// }

		// 	isSupplier, err := s.repo.ChackDuplicateSupplier(ctx, tx, modelSuppliers)
		// 	if err != nil {
		// 		return err
		// 	}

		// 	if isSupplier {
		// 		return nil
		// 	}

		// 	suppliersId, err := s.repo.InsertSupplier(ctx, tx, modelSuppliers)
		// 	if err != nil {
		// 		return err
		// 	}

		// 	response.SupplierID = suppliersId

		return nil
	})

	if err != nil {
		return response, err
	}

	return response, nil
}
