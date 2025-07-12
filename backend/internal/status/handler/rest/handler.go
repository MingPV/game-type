package rest

import (
	"github.com/MingPV/clean-go-template/internal/status/dto"
	"github.com/MingPV/clean-go-template/internal/status/usecase"

	// "github.com/MingPV/clean-go-template/internal/status/usecase"
	"github.com/MingPV/clean-go-template/internal/entities"
	responses "github.com/MingPV/clean-go-template/pkg/responses"
	"github.com/gofiber/fiber/v2"
)

type HttpStatusHandler struct {
	statusUseCase usecase.StatusUseCase
}

func NewHttpStatusHandler(useCase usecase.StatusUseCase) *HttpStatusHandler {
	return &HttpStatusHandler{statusUseCase: useCase}
}

// CreateStatus godoc
// @Summary Create a new status
// @Tags statuses
// @Accept json
// @Produce json
// @Param status body entities.Status true "Status payload"
// @Success 201 {object} entities.Status
// @Router /statuses [post]
func (h *HttpStatusHandler) CreateStatus(c *fiber.Ctx) error {
	var req dto.CreateStatusRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.Error(c, fiber.StatusBadRequest, "invalid request")
	}

	status := &entities.Status{
		CharacterID: req.CharacterID,
		StatusPoint: req.StatusPoint,
		Attack:      req.Attack,
		Defense:     req.Defense,
		HP:          req.HP,
		MP:          req.MP,
		Critical:    req.Critical,
	}

	if err := h.statusUseCase.CreateStatus(status); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToStatusResponse(status))
}

// FindAllStatuses godoc
// @Summary Get all statuses
// @Tags statuses
// @Produce json
// @Success 200 {array} entities.Status
// @Router /statuses [get]
func (h *HttpStatusHandler) FindAllStatuses(c *fiber.Ctx) error {
	statuses, err := h.statusUseCase.FindAllStatuses()
	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(dto.ToStatusResponseList(statuses))
}

// FindStatusByCharacterID godoc
// @Summary Get status by ID
// @Tags statuses
// @Produce json
// @Param character_id path int true "Status ID"
// @Success 200 {object} entities.Status
// @Router /statuses/{character_id} [get]
func (h *HttpStatusHandler) FindStatusByCharacterID(c *fiber.Ctx) error {
	character_id := c.Params("character_id")

	status, err := h.statusUseCase.FindStatusByCharacterID(character_id)
	if err != nil {
		return responses.Error(c, fiber.StatusNotFound, err.Error())
	}

	return c.JSON(dto.ToStatusResponse(status))
}

// PatchStatus godoc
// @Summary Update an status partially
// @Tags statuses
// @Accept json
// @Produce json
// @Param character_id path int true "Status ID"
// @Param status body entities.Status true "Status update payload"
// @Success 200 {object} entities.Status
// @Router /statuses/{character_id} [patch]
func (h *HttpStatusHandler) PatchStatus(c *fiber.Ctx) error {
	character_id := c.Params("character_id")

	var req dto.CreateStatusRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.Error(c, fiber.StatusBadRequest, "invalid request")
	}

	status := &entities.Status{
		StatusPoint: req.StatusPoint,
		Attack:      req.Attack,
		Defense:     req.Defense,
		HP:          req.HP,
		MP:          req.MP,
		Critical:    req.Critical,
	}
	if err := h.statusUseCase.PatchStatus(character_id, status); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	updatedStatus, err := h.statusUseCase.FindStatusByCharacterID(character_id)
	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(dto.ToStatusResponse(updatedStatus))
}

// DeleteStatus godoc
// @Summary Delete an status by ID
// @Tags statuses
// @Produce json
// @Param character_id path int true "Status ID"
// @Success 200 {object} response.MessageResponse
// @Router /statuses/{character_id} [delete]
func (h *HttpStatusHandler) DeleteStatus(c *fiber.Ctx) error {
	character_id := c.Params("character_id")

	if err := h.statusUseCase.DeleteStatus(character_id); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return responses.Message(c, fiber.StatusOK, "status deleted")
}
