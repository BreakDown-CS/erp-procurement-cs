package purchase_requests

import (
	"github.com/BreakDown-CS/erp-procurement-cs/modules/purchase_requests/handler"
	"github.com/gofiber/fiber/v2"
)

func purchaseRequestsRouter(app *fiber.App, handler *handler.Handler) {
	api := app.Group("/pr")

	api.Post("/savePurchaseRequests", handler.SavePurchaseRequests)

}
