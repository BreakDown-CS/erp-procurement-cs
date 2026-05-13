package posts

import (
	"context"

	"github.com/BreakDown-CS/erp-procurement-cs/modules/suppliers/dto"
)

type SuppliersService interface {
	CreateSupplier(ctx context.Context, supplier dto.CreateSuppliersRequest) (dto.StaffSuppliersResponse, error)
}
