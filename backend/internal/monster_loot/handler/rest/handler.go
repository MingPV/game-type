package rest

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/MingPV/clean-go-template/internal/monster_loot/dto"
	"github.com/MingPV/clean-go-template/internal/monster_loot/usecase"
	responses "github.com/MingPV/clean-go-template/pkg/responses"
	"github.com/gofiber/fiber/v2"
)

type HttpMonsterLootHandler struct {
	monsterLootUseCase usecase.MonsterLootUseCase
}

func NewHttpMonsterLootHandler(useCase usecase.MonsterLootUseCase) *HttpMonsterLootHandler {
	return &HttpMonsterLootHandler{monsterLootUseCase: useCase}
}

// CreateMonsterLoot godoc
// @Summary Create a new monsterLoot
// @Tags monsterLoots
// @Accept json
// @Produce json
// @Param monsterLoot body entities.MonsterLoot true "MonsterLoot payload"
// @Success 201 {object} entities.MonsterLoot
// @Router /monsterLoots [post]
func (h *HttpMonsterLootHandler) CreateMonsterLoot(c *fiber.Ctx) error {
	var req dto.CreateMonsterLootRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.Error(c, fiber.StatusBadRequest, "invalid request")
	}

	// monsterLoot := &entities.MonsterLoot{Total: req.Total}

	monsterLoot := &entities.MonsterLoot{
		MonsterID:   req.MonsterID,
		ItemID:      req.ItemID,
		DropRate:    req.DropRate,
		QuantityMin: req.QuantityMin,
		QuantityMax: req.QuantityMax,
	}

	if err := h.monsterLootUseCase.CreateMonsterLoot(monsterLoot); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToMonsterLootResponse(monsterLoot))
}

// FindAllMonsterLoots godoc
// @Summary Get all monsterLoots
// @Tags monsterLoots
// @Produce json
// @Success 200 {array} entities.MonsterLoot
// @Router /monsterLoots [get]
func (h *HttpMonsterLootHandler) FindAllMonsterLoots(c *fiber.Ctx) error {
	monsterLoots, err := h.monsterLootUseCase.FindAllMonsterLoots()
	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(dto.ToMonsterLootResponseList(monsterLoots))
}

// FindMonsterLootByID godoc
// @Summary Get monsterLoot by ID
// @Tags monsterLoots
// @Produce json
// @Param id path int true "MonsterLoot ID"
// @Success 200 {object} entities.MonsterLoot
// @Router /monsterLoots/{id} [get]
func (h *HttpMonsterLootHandler) FindMonsterLootByID(c *fiber.Ctx) error {
	id := c.Params("id")

	monsterLoot, err := h.monsterLootUseCase.FindMonsterLootByID(id)
	if err != nil {
		return responses.Error(c, fiber.StatusNotFound, err.Error())
	}

	return c.JSON(dto.ToMonsterLootResponse(monsterLoot))
}

// PatchMonsterLoot godoc
// @Summary Update an monsterLoot partially
// @Tags monsterLoots
// @Accept json
// @Produce json
// @Param id path int true "MonsterLoot ID"
// @Param monsterLoot body entities.MonsterLoot true "MonsterLoot update payload"
// @Success 200 {object} entities.MonsterLoot
// @Router /monsterLoots/{id} [patch]
func (h *HttpMonsterLootHandler) PatchMonsterLoot(c *fiber.Ctx) error {
	id := c.Params("id")

	var req dto.CreateMonsterLootRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.Error(c, fiber.StatusBadRequest, "invalid request")
	}

	monsterLoot := &entities.MonsterLoot{DropRate: req.DropRate}
	if err := h.monsterLootUseCase.PatchMonsterLoot(id, monsterLoot); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	updatedMonsterLoot, err := h.monsterLootUseCase.FindMonsterLootByID(id)
	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(dto.ToMonsterLootResponse(updatedMonsterLoot))
}

// DeleteMonsterLoot godoc
// @Summary Delete an monsterLoot by ID
// @Tags monsterLoots
// @Produce json
// @Param id path int true "MonsterLoot ID"
// @Success 200 {object} response.MessageResponse
// @Router /monsterLoots/{id} [delete]
func (h *HttpMonsterLootHandler) DeleteMonsterLoot(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.monsterLootUseCase.DeleteMonsterLoot(id); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return responses.Message(c, fiber.StatusOK, "monsterLoot deleted")
}
