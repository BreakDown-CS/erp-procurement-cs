package suppliers

import (
	"github.com/BreakDown-CS/erp-procurement-cs/modules/suppliers/handler"
	"github.com/BreakDown-CS/erp-procurement-cs/modules/suppliers/repository"
	"github.com/BreakDown-CS/erp-procurement-cs/modules/suppliers/service"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Wire(app *fiber.App, db *pgxpool.Pool) {
	repo := repository.NewRepository(db)
	service := service.NewService(repo, db)
	handler := handler.NewHandler(service)

	suppliersRouter(app, handler)
}
