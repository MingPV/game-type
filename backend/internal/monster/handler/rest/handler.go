package rest

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/MingPV/clean-go-template/internal/monster/dto"
	"github.com/MingPV/clean-go-template/internal/monster/usecase"
	responses "github.com/MingPV/clean-go-template/pkg/responses"
	"github.com/gofiber/fiber/v2"
)

type HttpMonsterHandler struct {
	monsterUseCase usecase.MonsterUseCase
}

func NewHttpMonsterHandler(useCase usecase.MonsterUseCase) *HttpMonsterHandler {
	return &HttpMonsterHandler{monsterUseCase: useCase}
}

// CreateMonster godoc
// @Summary Create a new monster
// @Tags monsters
// @Accept json
// @Produce json
// @Param monster body entities.Monster true "Monster payload"
// @Success 201 {object} entities.Monster
// @Router /monsters [post]
func (h *HttpMonsterHandler) CreateMonster(c *fiber.Ctx) error {
	var req dto.CreateMonsterRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.Error(c, fiber.StatusBadRequest, "invalid request")
	}

	// monster := &entities.Monster{Total: req.Total}

	monster := &entities.Monster{
		Name:          req.Name,
		Description:   req.Description,
		Level:         req.Level,
		HP:            req.HP,
		Attack:        req.Attack,
		Defense:       req.Defense,
		ExpReward:     req.ExpReward,
		GoldReward:    req.GoldReward,
		MonsterTypeID: req.MonsterTypeID,
	}

	if err := h.monsterUseCase.CreateMonster(monster); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToMonsterResponse(monster))
}

// FindAllMonsters godoc
// @Summary Get all monsters
// @Tags monsters
// @Produce json
// @Success 200 {array} entities.Monster
// @Router /monsters [get]
func (h *HttpMonsterHandler) FindAllMonsters(c *fiber.Ctx) error {
	monsters, err := h.monsterUseCase.FindAllMonsters()
	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(dto.ToMonsterResponseList(monsters))
}

// FindMonsterByID godoc
// @Summary Get monster by ID
// @Tags monsters
// @Produce json
// @Param id path int true "Monster ID"
// @Success 200 {object} entities.Monster
// @Router /monsters/{id} [get]
func (h *HttpMonsterHandler) FindMonsterByID(c *fiber.Ctx) error {
	id := c.Params("id")

	monster, err := h.monsterUseCase.FindMonsterByID(id)
	if err != nil {
		return responses.Error(c, fiber.StatusNotFound, err.Error())
	}

	return c.JSON(dto.ToMonsterResponse(monster))
}

// PatchMonster godoc
// @Summary Update an monster partially
// @Tags monsters
// @Accept json
// @Produce json
// @Param id path int true "Monster ID"
// @Param monster body entities.Monster true "Monster update payload"
// @Success 200 {object} entities.Monster
// @Router /monsters/{id} [patch]
func (h *HttpMonsterHandler) PatchMonster(c *fiber.Ctx) error {
	id := c.Params("id")

	var req dto.CreateMonsterRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.Error(c, fiber.StatusBadRequest, "invalid request")
	}

	monster := &entities.Monster{Description: req.Description}
	if err := h.monsterUseCase.PatchMonster(id, monster); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	updatedMonster, err := h.monsterUseCase.FindMonsterByID(id)
	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(dto.ToMonsterResponse(updatedMonster))
}

// DeleteMonster godoc
// @Summary Delete an monster by ID
// @Tags monsters
// @Produce json
// @Param id path int true "Monster ID"
// @Success 200 {object} response.MessageResponse
// @Router /monsters/{id} [delete]
func (h *HttpMonsterHandler) DeleteMonster(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.monsterUseCase.DeleteMonster(id); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return responses.Message(c, fiber.StatusOK, "monster deleted")
}
