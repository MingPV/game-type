package rest

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/MingPV/clean-go-template/internal/monster_type/dto"
	"github.com/MingPV/clean-go-template/internal/monster_type/usecase"
	responses "github.com/MingPV/clean-go-template/pkg/responses"
	"github.com/gofiber/fiber/v2"
)

type HttpMonsterTypeHandler struct {
	monsterTypeUseCase usecase.MonsterTypeUseCase
}

func NewHttpMonsterTypeHandler(useCase usecase.MonsterTypeUseCase) *HttpMonsterTypeHandler {
	return &HttpMonsterTypeHandler{monsterTypeUseCase: useCase}
}

// CreateMonsterType godoc
// @Summary Create a new monsterType
// @Tags monsterTypes
// @Accept json
// @Produce json
// @Param monsterType body entities.MonsterType true "MonsterType payload"
// @Success 201 {object} entities.MonsterType
// @Router /monsterTypes [post]
func (h *HttpMonsterTypeHandler) CreateMonsterType(c *fiber.Ctx) error {
	var req dto.CreateMonsterTypeRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.Error(c, fiber.StatusBadRequest, "invalid request")
	}

	// monsterType := &entities.MonsterType{Total: req.Total}

	monsterType := &entities.MonsterType{
		Name: req.Name,
	}

	if err := h.monsterTypeUseCase.CreateMonsterType(monsterType); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToMonsterTypeResponse(monsterType))
}

// FindAllMonsterTypes godoc
// @Summary Get all monsterTypes
// @Tags monsterTypes
// @Produce json
// @Success 200 {array} entities.MonsterType
// @Router /monsterTypes [get]
func (h *HttpMonsterTypeHandler) FindAllMonsterTypes(c *fiber.Ctx) error {
	monsterTypes, err := h.monsterTypeUseCase.FindAllMonsterTypes()
	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(dto.ToMonsterTypeResponseList(monsterTypes))
}

// FindMonsterTypeByID godoc
// @Summary Get monsterType by ID
// @Tags monsterTypes
// @Produce json
// @Param id path int true "MonsterType ID"
// @Success 200 {object} entities.MonsterType
// @Router /monsterTypes/{id} [get]
func (h *HttpMonsterTypeHandler) FindMonsterTypeByID(c *fiber.Ctx) error {
	id := c.Params("id")

	monsterType, err := h.monsterTypeUseCase.FindMonsterTypeByID(id)
	if err != nil {
		return responses.Error(c, fiber.StatusNotFound, err.Error())
	}

	return c.JSON(dto.ToMonsterTypeResponse(monsterType))
}

// PatchMonsterType godoc
// @Summary Update an monsterType partially
// @Tags monsterTypes
// @Accept json
// @Produce json
// @Param id path int true "MonsterType ID"
// @Param monsterType body entities.MonsterType true "MonsterType update payload"
// @Success 200 {object} entities.MonsterType
// @Router /monsterTypes/{id} [patch]
func (h *HttpMonsterTypeHandler) PatchMonsterType(c *fiber.Ctx) error {
	id := c.Params("id")

	var req dto.CreateMonsterTypeRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.Error(c, fiber.StatusBadRequest, "invalid request")
	}

	monsterType := &entities.MonsterType{Name: req.Name}
	if err := h.monsterTypeUseCase.PatchMonsterType(id, monsterType); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	updatedMonsterType, err := h.monsterTypeUseCase.FindMonsterTypeByID(id)
	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(dto.ToMonsterTypeResponse(updatedMonsterType))
}

// DeleteMonsterType godoc
// @Summary Delete an monsterType by ID
// @Tags monsterTypes
// @Produce json
// @Param id path int true "MonsterType ID"
// @Success 200 {object} response.MessageResponse
// @Router /monsterTypes/{id} [delete]
func (h *HttpMonsterTypeHandler) DeleteMonsterType(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.monsterTypeUseCase.DeleteMonsterType(id); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return responses.Message(c, fiber.StatusOK, "monsterType deleted")
}
