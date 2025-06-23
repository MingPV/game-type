package rest

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/MingPV/clean-go-template/internal/inventory/dto"
	"github.com/MingPV/clean-go-template/internal/inventory/usecase"
	responses "github.com/MingPV/clean-go-template/pkg/responses"
	"github.com/gofiber/fiber/v2"
)

type HttpInventoryHandler struct {
	inventoryUseCase usecase.InventoryUseCase
}

func NewHttpInventoryHandler(useCase usecase.InventoryUseCase) *HttpInventoryHandler {
	return &HttpInventoryHandler{inventoryUseCase: useCase}
}

// CreateInventory godoc
// @Summary Create a new inventory
// @Tags inventories
// @Accept json
// @Produce json
// @Param inventory body entities.Inventory true "Inventory payload"
// @Success 201 {object} entities.Inventory
// @Router /inventories [post]
func (h *HttpInventoryHandler) CreateInventory(c *fiber.Ctx) error {
	var req dto.CreateInventoryRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.Error(c, fiber.StatusBadRequest, "invalid request")
	}

	inventory := &entities.Inventory{
		MaxSlots: req.MaxSlots,
	}

	if err := h.inventoryUseCase.CreateInventory(inventory); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToInventoryResponse(inventory))
}

// FindAllInventories godoc
// @Summary Get all inventories
// @Tags inventories
// @Produce json
// @Success 200 {array} entities.Inventory
// @Router /inventories [get]
func (h *HttpInventoryHandler) FindAllInventories(c *fiber.Ctx) error {
	inventories, err := h.inventoryUseCase.FindAllInventories()
	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(dto.ToInventoryResponseList(inventories))
}

// FindInventoryByID godoc
// @Summary Get inventory by ID
// @Tags inventories
// @Produce json
// @Param id path int true "Inventory ID"
// @Success 200 {object} entities.Inventory
// @Router /inventories/{id} [get]
func (h *HttpInventoryHandler) FindInventoryByID(c *fiber.Ctx) error {
	id := c.Params("id")

	inventory, err := h.inventoryUseCase.FindInventoryByID(id)
	if err != nil {
		return responses.Error(c, fiber.StatusNotFound, err.Error())
	}

	return c.JSON(dto.ToInventoryResponse(inventory))
}

// PatchInventory godoc
// @Summary Update an inventory partially
// @Tags inventories
// @Accept json
// @Produce json
// @Param id path int true "Inventory ID"
// @Param inventory body entities.Inventory true "Inventory update payload"
// @Success 200 {object} entities.Inventory
// @Router /inventories/{id} [patch]
func (h *HttpInventoryHandler) PatchInventory(c *fiber.Ctx) error {
	id := c.Params("id")

	var req dto.CreateInventoryRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.Error(c, fiber.StatusBadRequest, "invalid request")
	}

	inventory := &entities.Inventory{MaxSlots: req.MaxSlots}
	if err := h.inventoryUseCase.PatchInventory(id, inventory); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	updatedInventory, err := h.inventoryUseCase.FindInventoryByID(id)
	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(dto.ToInventoryResponse(updatedInventory))
}

// DeleteInventory godoc
// @Summary Delete an inventory by ID
// @Tags inventories
// @Produce json
// @Param id path int true "Inventory ID"
// @Success 200 {object} response.MessageResponse
// @Router /inventories/{id} [delete]
func (h *HttpInventoryHandler) DeleteInventory(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.inventoryUseCase.DeleteInventory(id); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return responses.Message(c, fiber.StatusOK, "inventory deleted")
}
