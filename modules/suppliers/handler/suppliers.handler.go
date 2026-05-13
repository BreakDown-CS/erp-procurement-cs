package handler

import (
	"github.com/BreakDown-CS/erp-procurement-cs/internal/helper"
	"github.com/BreakDown-CS/erp-procurement-cs/modules/suppliers/dto"
	ports "github.com/BreakDown-CS/erp-procurement-cs/modules/suppliers/posts"
	"github.com/BreakDown-CS/erp-procurement-cs/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Handler struct {
	service ports.SuppliersService
}

func NewHandler(service ports.SuppliersService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) SaveSupplier(c *fiber.Ctx) error {
	req := dto.CreateSuppliersRequest{}

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

	result, err := h.service.CreateSupplier(ctx, req)
	if err != nil {
		return response.Error(c, 500, err)
	}

	if result.SupplierID == uuid.Nil {
		return response.Warning(c, 400, "supplier code already exist")
	}

	return response.Created(c, result, "save supplier success")
}
