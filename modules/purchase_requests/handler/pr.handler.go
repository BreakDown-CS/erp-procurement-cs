package handler

import (
	"github.com/BreakDown-CS/erp-procurement-cs/internal/helper"
	"github.com/BreakDown-CS/erp-procurement-cs/modules/purchase_requests/dto"
	ports "github.com/BreakDown-CS/erp-procurement-cs/modules/purchase_requests/posts"
	"github.com/BreakDown-CS/erp-procurement-cs/response"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service ports.PurchaseRequestService
}

func NewHandler(service ports.PurchaseRequestService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) SavePurchaseRequests(c *fiber.Ctx) error {
	req := dto.CreatPurchaseRequest{}

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if errors := helper.ValidateStruct(req); errors != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errors,
		})
	}

	ctx := c.Context()

	result, err := h.service.CreatePurchaseRequests(ctx, req)
	if err != nil {
		return response.Error(c, 500, err)
	}

	if result.PRNo == "" {
		return response.Warning(c, 400, "supplier code already exist")
	}

	return response.Created(c, result, "save supplier success")
}
