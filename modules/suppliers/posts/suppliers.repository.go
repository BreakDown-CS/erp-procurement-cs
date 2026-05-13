package posts

import (
	"context"

	"github.com/BreakDown-CS/erp-procurement-cs/modules/suppliers/model"
	"github.com/jackc/pgx/v5"
)

type SuppliersRepository interface {
	ChackDuplicateSupplier(ctx context.Context, tx pgx.Tx, supplier model.Suppliers) (bool, error)
}
