package rest

import (
	"github.com/MingPV/clean-go-template/internal/level_progress/dto"
	"github.com/MingPV/clean-go-template/internal/level_progress/usecase"

	// "github.com/MingPV/clean-go-template/internal/level_progress/usecase"
	"github.com/MingPV/clean-go-template/internal/entities"
	responses "github.com/MingPV/clean-go-template/pkg/responses"
	"github.com/gofiber/fiber/v2"
)

type HttpLevelProgressHandler struct {
	level_progressUseCase usecase.LevelProgressUseCase
}

func NewHttpLevelProgressHandler(useCase usecase.LevelProgressUseCase) *HttpLevelProgressHandler {
	return &HttpLevelProgressHandler{level_progressUseCase: useCase}
}

// CreateLevelProgress godoc
// @Summary Create a new level_progress
// @Tags level_progresses
// @Accept json
// @Produce json
// @Param level_progress body entities.LevelProgress true "LevelProgress payload"
// @Success 201 {object} entities.LevelProgress
// @Router /level_progresses [post]
func (h *HttpLevelProgressHandler) CreateLevelProgress(c *fiber.Ctx) error {
	var req dto.CreateLevelProgressRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.Error(c, fiber.StatusBadRequest, "invalid request")
	}

	level_progress := &entities.LevelProgress{
		Level:       req.Level,
		ExpRequired: req.ExpRequired,
	}

	if err := h.level_progressUseCase.CreateLevelProgress(level_progress); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToLevelProgressResponse(level_progress))
}

// FindAllLevelProgresses godoc
// @Summary Get all level_progresses
// @Tags level_progresses
// @Produce json
// @Success 200 {array} entities.LevelProgress
// @Router /level_progresses [get]
func (h *HttpLevelProgressHandler) FindAllLevelProgresses(c *fiber.Ctx) error {
	level_progresses, err := h.level_progressUseCase.FindAllLevelProgresses()
	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(dto.ToLevelProgressResponseList(level_progresses))
}

// FindLevelProgressByID godoc
// @Summary Get level_progress by ID
// @Tags level_progresses
// @Produce json
// @Param id path int true "LevelProgress ID"
// @Success 200 {object} entities.LevelProgress
// @Router /level_progresses/{id} [get]
func (h *HttpLevelProgressHandler) FindLevelProgressByID(c *fiber.Ctx) error {
	id := c.Params("id")

	level_progress, err := h.level_progressUseCase.FindLevelProgressByID(id)
	if err != nil {
		return responses.Error(c, fiber.StatusNotFound, err.Error())
	}

	return c.JSON(dto.ToLevelProgressResponse(level_progress))
}

// PatchLevelProgress godoc
// @Summary Update an level_progress partially
// @Tags level_progresses
// @Accept json
// @Produce json
// @Param id path int true "LevelProgress ID"
// @Param level_progress body entities.LevelProgress true "LevelProgress update payload"
// @Success 200 {object} entities.LevelProgress
// @Router /level_progresses/{id} [patch]
func (h *HttpLevelProgressHandler) PatchLevelProgress(c *fiber.Ctx) error {
	id := c.Params("id")

	var req dto.CreateLevelProgressRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.Error(c, fiber.StatusBadRequest, "invalid request")
	}

	level_progress := &entities.LevelProgress{ExpRequired: req.ExpRequired}
	if err := h.level_progressUseCase.PatchLevelProgress(id, level_progress); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	updatedLevelProgress, err := h.level_progressUseCase.FindLevelProgressByID(id)
	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(dto.ToLevelProgressResponse(updatedLevelProgress))
}

// DeleteLevelProgress godoc
// @Summary Delete an level_progress by ID
// @Tags level_progresses
// @Produce json
// @Param id path int true "LevelProgress ID"
// @Success 200 {object} response.MessageResponse
// @Router /level_progresses/{id} [delete]
func (h *HttpLevelProgressHandler) DeleteLevelProgress(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.level_progressUseCase.DeleteLevelProgress(id); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return responses.Message(c, fiber.StatusOK, "level_progress deleted")
}
