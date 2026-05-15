package service

import (
	"context"

	"github.com/BreakDown-CS/erp-procurement-cs/internal/helper"
	"github.com/BreakDown-CS/erp-procurement-cs/modules/purchase_requests/dto"
	"github.com/BreakDown-CS/erp-procurement-cs/modules/purchase_requests/model"
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

func (s *service) CreatePurchaseRequests(ctx context.Context, pr dto.CreatPurchaseRequest) (dto.PurchaseRequestResponse, error) {

	response := dto.PurchaseRequestResponse{}

	err := s.helper.WithTx(ctx, func(tx pgx.Tx) error {

		prNoNew, err := s.repo.GeneratePRNO(ctx, tx)
		if err != nil {
			return err
		}

		RequestDateConvert := helper.MustConvertStringToTime(pr.RequestDate)

		modelPurchaseRequests := model.PurchaseRequests{
			PRNo:           prNoNew,
			RequestDate:    RequestDateConvert,
			StaffRequestId: pr.StaffId,
			DepartmentID:   pr.DepartmentId,
			CreatedBy:      pr.StaffId,
		}

		prId, err := s.repo.InsertPurchaseRequest(ctx, tx, modelPurchaseRequests)
		if err != nil {
			return err
		}

		response.PRId = prId

		for _, prDetail := range pr.Items {

			modelPurchaseRequestsDetail := model.PurchaseRequestDetails{
				PurchaseRequestID: prId,
				ProductId:         prDetail.ProductID,
				Qty:               float64(prDetail.Qty),
				UnitPrice:         float64(prDetail.UnitPrice),
				TotalPrice:        float64(prDetail.UnitPrice * prDetail.Qty),
				CreatedBy:         pr.StaffId,
			}

			prDetailId, err := s.repo.InsertPurchaseRequestDetail(ctx, tx, modelPurchaseRequestsDetail)
			if err != nil {
				return err
			}

			response.PRDetailId = append(response.PRDetailId, prDetailId)

			modelPurchaseRequestsApprove := model.PurchaseRequestApproved{
				PurchaseRequestDetailID: prDetailId,
				UpdatedBy:               pr.StaffId,
			}

			prApprovedId, err := s.repo.InsertPurchaseRequestApproved(ctx, tx, modelPurchaseRequestsApprove)
			if err != nil {
				return err
			}

			response.PRApprovedId = append(response.PRApprovedId, prApprovedId)
		}

		return nil
	})

	if err != nil {
		return response, err
	}

	return response, nil
}
