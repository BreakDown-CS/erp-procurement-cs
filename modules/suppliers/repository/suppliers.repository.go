package repository

import (
	"context"

	"github.com/BreakDown-CS/erp-procurement-cs/modules/suppliers/model"
	ports "github.com/BreakDown-CS/erp-procurement-cs/modules/suppliers/posts"
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

	queryChackSupplier := `SELECT EXISTS(SELECT 1 FROM suppliers WHERE supplier_code = $1)`
	if err := tx.QueryRow(ctx, queryChackSupplier, supplier.SupplierCode).Scan(&isSuppliers); err != nil && err != pgx.ErrNoRows {
		return false, err
	}

	return isSuppliers, nil
}
