package rest

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/MingPV/clean-go-template/internal/item_instance/dto"
	"github.com/MingPV/clean-go-template/internal/item_instance/usecase"
	responses "github.com/MingPV/clean-go-template/pkg/responses"
	"github.com/gofiber/fiber/v2"
)

type HttpItemInstanceHandler struct {
	itemInstanceUseCase usecase.ItemInstanceUseCase
}

func NewHttpItemInstanceHandler(useCase usecase.ItemInstanceUseCase) *HttpItemInstanceHandler {
	return &HttpItemInstanceHandler{itemInstanceUseCase: useCase}
}

// CreateItemInstance godoc
// @Summary Create a new itemInstance
// @Tags itemInstances
// @Accept json
// @Produce json
// @Param itemInstance body entities.ItemInstance true "ItemInstance payload"
// @Success 201 {object} entities.ItemInstance
// @Router /itemInstances [post]
func (h *HttpItemInstanceHandler) CreateItemInstance(c *fiber.Ctx) error {
	var req dto.CreateItemInstanceRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.Error(c, fiber.StatusBadRequest, "invalid request")
	}

	itemInstance := &entities.ItemInstance{
		InventoryID:      req.InventoryID,
		ItemID:           req.ItemID,
		UpgradeLevel:     req.UpgradeLevel,
		OwnerCharacterID: req.OwnerCharacterID,
	}

	if err := h.itemInstanceUseCase.CreateItemInstance(itemInstance); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToItemInstanceResponse(itemInstance))
}

// FindAllItemInstances godoc
// @Summary Get all itemInstances
// @Tags itemInstances
// @Produce json
// @Success 200 {array} entities.ItemInstance
// @Router /itemInstances [get]
func (h *HttpItemInstanceHandler) FindAllItemInstances(c *fiber.Ctx) error {
	itemInstances, err := h.itemInstanceUseCase.FindAllItemInstances()
	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(dto.ToItemInstanceResponseList(itemInstances))
}

// FindItemInstanceByID godoc
// @Summary Get itemInstance by ID
// @Tags itemInstances
// @Produce json
// @Param id path int true "ItemInstance ID"
// @Success 200 {object} entities.ItemInstance
// @Router /itemInstances/{id} [get]
func (h *HttpItemInstanceHandler) FindItemInstanceByID(c *fiber.Ctx) error {
	id := c.Params("id")

	itemInstance, err := h.itemInstanceUseCase.FindItemInstanceByID(id)
	if err != nil {
		return responses.Error(c, fiber.StatusNotFound, err.Error())
	}

	return c.JSON(dto.ToItemInstanceResponse(itemInstance))
}

// PatchItemInstance godoc
// @Summary Update an itemInstance partially
// @Tags itemInstances
// @Accept json
// @Produce json
// @Param id path int true "ItemInstance ID"
// @Param itemInstance body entities.ItemInstance true "ItemInstance update payload"
// @Success 200 {object} entities.ItemInstance
// @Router /itemInstances/{id} [patch]
// func (h *HttpItemInstanceHandler) PatchItemInstance(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	itemInstanceID, err := strconv.Atoi(id)
// 	if err != nil {
// 		return responses.Error(c, fiber.StatusBadRequest, "invalid id")
// 	}

// 	var req dto.CreateItemInstanceRequest
// 	if err := c.BodyParser(&req); err != nil {
// 		return responses.Error(c, fiber.StatusBadRequest, "invalid request")
// 	}

// 	itemInstance := &entities.ItemInstance{Total: req.Total}
// 	if err := h.itemInstanceUseCase.PatchItemInstance(itemInstanceID, itemInstance); err != nil {
// 		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
// 	}

// 	updatedItemInstance, err := h.itemInstanceUseCase.FindItemInstanceByID(itemInstanceID)
// 	if err != nil {
// 		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
// 	}

// 	return c.JSON(dto.ToItemInstanceResponse(updatedItemInstance))
// }

// DeleteItemInstance godoc
// @Summary Delete an itemInstance by ID
// @Tags itemInstances
// @Produce json
// @Param id path int true "ItemInstance ID"
// @Success 200 {object} response.MessageResponse
// @Router /itemInstances/{id} [delete]
func (h *HttpItemInstanceHandler) DeleteItemInstance(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.itemInstanceUseCase.DeleteItemInstance(id); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return responses.Message(c, fiber.StatusOK, "itemInstance deleted")
}
