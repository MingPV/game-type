package rest

import (
	"encoding/json"

	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/MingPV/clean-go-template/internal/item_level_stat/dto"
	"github.com/MingPV/clean-go-template/internal/item_level_stat/usecase"
	responses "github.com/MingPV/clean-go-template/pkg/responses"
	"github.com/gofiber/fiber/v2"
)

type HttpItemLevelStatHandler struct {
	itemLevelStatUseCase usecase.ItemLevelStatUseCase
}

func NewHttpItemLevelStatHandler(useCase usecase.ItemLevelStatUseCase) *HttpItemLevelStatHandler {
	return &HttpItemLevelStatHandler{itemLevelStatUseCase: useCase}
}

// CreateItemLevelStat godoc
// @Summary Create a new itemLevelStat
// @Tags itemLevelStats
// @Accept json
// @Produce json
// @Param itemLevelStat body entities.ItemLevelStat true "ItemLevelStat payload"
// @Success 201 {object} entities.ItemLevelStat
// @Router /itemLevelStats [post]
func (h *HttpItemLevelStatHandler) CreateItemLevelStat(c *fiber.Ctx) error {
	var req dto.CreateItemLevelStatRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.Error(c, fiber.StatusBadRequest, "invalid request")
	}

	// convert map to string
	bonusStatBytes, err := json.Marshal(req.BonusStat)
	if err != nil {
		return responses.Error(c, fiber.StatusBadRequest, "invalid bonus_stat format")
	}

	itemLevelStat := &entities.ItemLevelStat{
		ItemID:    req.ItemID,
		BonusStat: string(bonusStatBytes),
	}

	if err := h.itemLevelStatUseCase.CreateItemLevelStat(itemLevelStat); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToItemLevelStatResponse(itemLevelStat))
}

// FindAllItemLevelStats godoc
// @Summary Get all itemLevelStats
// @Tags itemLevelStats
// @Produce json
// @Success 200 {array} entities.ItemLevelStat
// @Router /itemLevelStats [get]
func (h *HttpItemLevelStatHandler) FindAllItemLevelStats(c *fiber.Ctx) error {
	itemLevelStats, err := h.itemLevelStatUseCase.FindAllItemLevelStats()
	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(dto.ToItemLevelStatResponseList(itemLevelStats))
}

// FindItemLevelStatByID godoc
// @Summary Get itemLevelStat by ID
// @Tags itemLevelStats
// @Produce json
// @Param id path int true "ItemLevelStat ID"
// @Success 200 {object} entities.ItemLevelStat
// @Router /itemLevelStats/{id} [get]
func (h *HttpItemLevelStatHandler) FindItemLevelStatByID(c *fiber.Ctx) error {
	id := c.Params("id")

	itemLevelStat, err := h.itemLevelStatUseCase.FindItemLevelStatByID(id)
	if err != nil {
		return responses.Error(c, fiber.StatusNotFound, err.Error())
	}

	return c.JSON(dto.ToItemLevelStatResponse(itemLevelStat))
}

// PatchItemLevelStat godoc
// @Summary Update an itemLevelStat partially
// @Tags itemLevelStats
// @Accept json
// @Produce json
// @Param id path int true "ItemLevelStat ID"
// @Param itemLevelStat body entities.ItemLevelStat true "ItemLevelStat update payload"
// @Success 200 {object} entities.ItemLevelStat
// @Router /itemLevelStats/{id} [patch]
func (h *HttpItemLevelStatHandler) PatchItemLevelStat(c *fiber.Ctx) error {
	id := c.Params("id")

	var req dto.PatchItemLevelStatRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.Error(c, fiber.StatusBadRequest, "invalid request")
	}

	// convert map to string
	bonusStatBytes, err := json.Marshal(req.BonusStat)
	if err != nil {
		return responses.Error(c, fiber.StatusBadRequest, "invalid bonus_stat format")
	}

	itemLevelStat := &entities.ItemLevelStat{BonusStat: string(bonusStatBytes)}
	if err := h.itemLevelStatUseCase.PatchItemLevelStat(id, itemLevelStat); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	updatedItemLevelStat, err := h.itemLevelStatUseCase.FindItemLevelStatByID(id)
	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(dto.ToItemLevelStatResponse(updatedItemLevelStat))
}

// DeleteItemLevelStat godoc
// @Summary Delete an itemLevelStat by ID
// @Tags itemLevelStats
// @Produce json
// @Param id path int true "ItemLevelStat ID"
// @Success 200 {object} response.MessageResponse
// @Router /itemLevelStats/{id} [delete]
func (h *HttpItemLevelStatHandler) DeleteItemLevelStat(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.itemLevelStatUseCase.DeleteItemLevelStat(id); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return responses.Message(c, fiber.StatusOK, "itemLevelStat deleted")
}
