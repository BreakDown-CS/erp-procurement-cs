package repository

import (
	"context"

	"github.com/BreakDown-CS/erp-procurement-cs/modules/suppliers/model"
	ports "github.com/BreakDown-CS/erp-procurement-cs/modules/suppliers/posts"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) ports.SuppliersRepository {
	return &repository{db: db}
}

func (r *repository) ChackDuplicateSupplier(ctx context.Context, tx pgx.Tx, supplier model.Suppliers) (bool, error) {
	var isSuppliers bool

	queryChackSupplier := `SELECT EXISTS(SELECT 1 FROM erp.suppliers WHERE supplier_code = $1)`

	err := tx.QueryRow(ctx, queryChackSupplier, supplier.SupplierCode).Scan(&isSuppliers)

	if err != nil {
		return false, err
	}

	return isSuppliers, nil
}

func (r *repository) InsertSupplier(ctx context.Context, tx pgx.Tx, supplier model.Suppliers) (uuid.UUID, error) {
	var suppliersId uuid.UUID

	querySupplierInsert := `
		INSERT INTO erp.suppliers (
			supplier_code,
			supplier_name,
			tax_id,
			email,
			phone,
			address,
			created_by
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7
		) RETURNING id
	`

	err := tx.QueryRow(ctx,
		querySupplierInsert,
		supplier.SupplierCode,
		supplier.SupplierName,
		supplier.TaxID,
		supplier.Email,
		supplier.Phone,
		supplier.Address,
		supplier.CreatedBy,
	).Scan(&suppliersId)

	if err != nil {
		return uuid.Nil, err
	}

	return suppliersId, nil
}
