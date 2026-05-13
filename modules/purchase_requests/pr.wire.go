package purchase_requests

import (
	"github.com/BreakDown-CS/erp-procurement-cs/modules/purchase_requests/handler"
	"github.com/BreakDown-CS/erp-procurement-cs/modules/purchase_requests/repository"
	"github.com/BreakDown-CS/erp-procurement-cs/modules/purchase_requests/service"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Wire(app *fiber.App, db *pgxpool.Pool) {
	repo := repository.NewRepository(db)
	service := service.NewService(repo, db)
	handler := handler.NewHandler(service)

	purchaseRequestsRouter(app, handler)
}
