package service

import (
	"context"

	"github.com/BreakDown-CS/erp-procurement-cs/internal/helper"
	"github.com/BreakDown-CS/erp-procurement-cs/modules/suppliers/dto"
	"github.com/BreakDown-CS/erp-procurement-cs/modules/suppliers/model"
	ports "github.com/BreakDown-CS/erp-procurement-cs/modules/suppliers/posts"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type service struct {
	repo   ports.SuppliersRepository
	helper *helper.Uow
	db     *pgxpool.Pool
}

func NewService(repo ports.SuppliersRepository, db *pgxpool.Pool) ports.SuppliersService {
	return &service{
		repo:   repo,
		helper: helper.New(db),
		db:     db,
	}
}

func (s *service) CreateSupplier(ctx context.Context, supplier dto.CreateSuppliersRequest) (dto.StaffSuppliersResponse, error) {

	response := dto.StaffSuppliersResponse{}

	err := s.helper.WithTx(ctx, func(tx pgx.Tx) error {

		modelIsSuppliers := model.Suppliers{
			SupplierCode: supplier.SupplierCode,
		}

		isSupplier, err := s.repo.ChackDuplicateSupplier(ctx, tx, modelIsSuppliers)
		if err != nil {
			return err
		}

		if isSupplier {
			return nil
		}

		return nil
	})

	if err != nil {
		return response, err
	}
	return response, nil
}
