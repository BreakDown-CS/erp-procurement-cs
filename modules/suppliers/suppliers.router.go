package suppliers

import (
	"github.com/BreakDown-CS/erp-procurement-cs/modules/suppliers/handler"
	"github.com/gofiber/fiber/v2"
)

func suppliersRouter(app *fiber.App, handler *handler.Handler) {
	api := app.Group("/suppliers")

	api.Post("/saveSuppliers", handler.SaveSupplier)

}
