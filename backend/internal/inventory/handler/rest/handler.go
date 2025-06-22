package rest

import (
	"strconv"

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
// @Tags inventorys
// @Accept json
// @Produce json
// @Param inventory body entities.Inventory true "Inventory payload"
// @Success 201 {object} entities.Inventory
// @Router /inventorys [post]
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

// FindAllInventorys godoc
// @Summary Get all inventorys
// @Tags inventorys
// @Produce json
// @Success 200 {array} entities.Inventory
// @Router /inventorys [get]
func (h *HttpInventoryHandler) FindAllInventorys(c *fiber.Ctx) error {
	inventorys, err := h.inventoryUseCase.FindAllInventorys()
	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(dto.ToInventoryResponseList(inventorys))
}

// FindInventoryByID godoc
// @Summary Get inventory by ID
// @Tags inventorys
// @Produce json
// @Param id path int true "Inventory ID"
// @Success 200 {object} entities.Inventory
// @Router /inventorys/{id} [get]
func (h *HttpInventoryHandler) FindInventoryByID(c *fiber.Ctx) error {
	id := c.Params("id")
	inventoryID, err := strconv.Atoi(id)
	if err != nil {
		return responses.Error(c, fiber.StatusBadRequest, "invalid id")
	}

	inventory, err := h.inventoryUseCase.FindInventoryByID(inventoryID)
	if err != nil {
		return responses.Error(c, fiber.StatusNotFound, err.Error())
	}

	return c.JSON(dto.ToInventoryResponse(inventory))
}

// PatchInventory godoc
// @Summary Update an inventory partially
// @Tags inventorys
// @Accept json
// @Produce json
// @Param id path int true "Inventory ID"
// @Param inventory body entities.Inventory true "Inventory update payload"
// @Success 200 {object} entities.Inventory
// @Router /inventorys/{id} [patch]
// func (h *HttpInventoryHandler) PatchInventory(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	inventoryID, err := strconv.Atoi(id)
// 	if err != nil {
// 		return responses.Error(c, fiber.StatusBadRequest, "invalid id")
// 	}

// 	var req dto.CreateInventoryRequest
// 	if err := c.BodyParser(&req); err != nil {
// 		return responses.Error(c, fiber.StatusBadRequest, "invalid request")
// 	}

// 	inventory := &entities.Inventory{Total: req.Total}
// 	if err := h.inventoryUseCase.PatchInventory(inventoryID, inventory); err != nil {
// 		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
// 	}

// 	updatedInventory, err := h.inventoryUseCase.FindInventoryByID(inventoryID)
// 	if err != nil {
// 		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
// 	}

// 	return c.JSON(dto.ToInventoryResponse(updatedInventory))
// }

// DeleteInventory godoc
// @Summary Delete an inventory by ID
// @Tags inventorys
// @Produce json
// @Param id path int true "Inventory ID"
// @Success 200 {object} response.MessageResponse
// @Router /inventorys/{id} [delete]
func (h *HttpInventoryHandler) DeleteInventory(c *fiber.Ctx) error {
	id := c.Params("id")
	inventoryID, err := strconv.Atoi(id)
	if err != nil {
		return responses.Error(c, fiber.StatusBadRequest, "invalid id")
	}

	if err := h.inventoryUseCase.DeleteInventory(inventoryID); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return responses.Message(c, fiber.StatusOK, "inventory deleted")
}
