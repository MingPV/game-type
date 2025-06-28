package rest

import (
	"encoding/json"

	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/MingPV/clean-go-template/internal/item/dto"
	"github.com/MingPV/clean-go-template/internal/item/usecase"
	responses "github.com/MingPV/clean-go-template/pkg/responses"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type HttpItemHandler struct {
	itemUseCase usecase.ItemUseCase
}

func NewHttpItemHandler(useCase usecase.ItemUseCase) *HttpItemHandler {
	return &HttpItemHandler{itemUseCase: useCase}
}

// CreateItem godoc
// @Summary Create a new item
// @Tags items
// @Accept json
// @Produce json
// @Param item body entities.Item true "Item payload"
// @Success 201 {object} entities.Item
// @Router /items [post]
func (h *HttpItemHandler) CreateItem(c *fiber.Ctx) error {
	var req dto.CreateItemRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.Error(c, fiber.StatusBadRequest, "invalid request")
	}

	// use in item, item_level_stat
	item_id := uuid.New()

	item := &entities.Item{
		ID:            item_id,
		Name:          req.Name,
		Description:   req.Description,
		ItemTypeID:    req.ItemTypeID,
		RarityID:      req.RarityID,
		RequiredLevel: req.RequiredLevel,
		MaxStack:      req.MaxStack,
	}

	// convert map to string
	statBytes, err := json.Marshal(req.LevelStat)
	if err != nil {
		return responses.Error(c, fiber.StatusBadRequest, "invalid level_stat format")
	}

	item_level_stat := &entities.ItemLevelStat{
		ItemID:    item_id,
		BonusStat: string(statBytes),
	}

	return_item, err := h.itemUseCase.CreateItem(item, item_level_stat)

	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToItemResponse(return_item))
}

// FindAllItems godoc
// @Summary Get all items
// @Tags items
// @Produce json
// @Success 200 {array} entities.Item
// @Router /items [get]
func (h *HttpItemHandler) FindAllItems(c *fiber.Ctx) error {
	items, err := h.itemUseCase.FindAllItems()
	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(dto.ToItemResponseList(items))
}

// FindItemByID godoc
// @Summary Get item by ID
// @Tags items
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} entities.Item
// @Router /items/{id} [get]
func (h *HttpItemHandler) FindItemByID(c *fiber.Ctx) error {
	id := c.Params("id")

	item, err := h.itemUseCase.FindItemByID(id)
	if err != nil {
		return responses.Error(c, fiber.StatusNotFound, err.Error())
	}

	return c.JSON(dto.ToItemResponse(item))
}

// PatchItem godoc
// @Summary Update an item partially
// @Tags items
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param item body entities.Item true "Item update payload"
// @Success 200 {object} entities.Item
// @Router /items/{id} [patch]
func (h *HttpItemHandler) PatchItem(c *fiber.Ctx) error {
	id := c.Params("id")

	var req dto.CreateItemRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.Error(c, fiber.StatusBadRequest, "invalid request")
	}

	item := &entities.Item{
		Name:          req.Name,
		Description:   req.Description,
		ItemTypeID:    req.ItemTypeID,
		RarityID:      req.RarityID,
		RequiredLevel: req.RequiredLevel,
		MaxStack:      req.MaxStack,
	}
	if err := h.itemUseCase.PatchItem(id, item); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	updatedItem, err := h.itemUseCase.FindItemByID(id)
	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(dto.ToItemResponse(updatedItem))
}

// DeleteItem godoc
// @Summary Delete an item by ID
// @Tags items
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} response.MessageResponse
// @Router /items/{id} [delete]
func (h *HttpItemHandler) DeleteItem(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.itemUseCase.DeleteItem(id); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return responses.Message(c, fiber.StatusOK, "item deleted")
}
