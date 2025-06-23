package rest

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/MingPV/clean-go-template/internal/setting/dto"
	"github.com/MingPV/clean-go-template/internal/setting/usecase"
	responses "github.com/MingPV/clean-go-template/pkg/responses"
	"github.com/gofiber/fiber/v2"
)

type HttpSettingHandler struct {
	settingUseCase usecase.SettingUseCase
}

func NewHttpSettingHandler(useCase usecase.SettingUseCase) *HttpSettingHandler {
	return &HttpSettingHandler{settingUseCase: useCase}
}

// CreateSetting godoc
// @Summary Create a new setting
// @Tags settings
// @Accept json
// @Produce json
// @Param setting body entities.Setting true "Setting payload"
// @Success 201 {object} entities.Setting
// @Router /settings [post]
func (h *HttpSettingHandler) CreateSetting(c *fiber.Ctx) error {
	var req dto.CreateSettingRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.Error(c, fiber.StatusBadRequest, "invalid request")
	}

	// setting := &entities.Setting{Total: req.Total}

	setting := &entities.Setting{
		UserID:      req.UserID,
		MusicVolume: req.MusicVolume,
		Language:    req.Language,
	}

	if err := h.settingUseCase.CreateSetting(setting); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToSettingResponse(setting))
}

// FindAllSettings godoc
// @Summary Get all settings
// @Tags settings
// @Produce json
// @Success 200 {array} entities.Setting
// @Router /settings [get]
func (h *HttpSettingHandler) FindAllSettings(c *fiber.Ctx) error {
	settings, err := h.settingUseCase.FindAllSettings()
	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(dto.ToSettingResponseList(settings))
}

// FindSettingByID godoc
// @Summary Get setting by ID
// @Tags settings
// @Produce json
// @Param id path int true "Setting ID"
// @Success 200 {object} entities.Setting
// @Router /settings/{id} [get]
func (h *HttpSettingHandler) FindSettingByID(c *fiber.Ctx) error {
	id := c.Params("id")

	setting, err := h.settingUseCase.FindSettingByID(id)
	if err != nil {
		return responses.Error(c, fiber.StatusNotFound, err.Error())
	}

	return c.JSON(dto.ToSettingResponse(setting))
}

// PatchSetting godoc
// @Summary Update an setting partially
// @Tags settings
// @Accept json
// @Produce json
// @Param id path int true "Setting ID"
// @Param setting body entities.Setting true "Setting update payload"
// @Success 200 {object} entities.Setting
// @Router /settings/{id} [patch]
func (h *HttpSettingHandler) PatchSetting(c *fiber.Ctx) error {
	id := c.Params("id")

	var req dto.CreateSettingRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.Error(c, fiber.StatusBadRequest, "invalid request")
	}

	setting := &entities.Setting{Language: req.Language}
	if err := h.settingUseCase.PatchSetting(id, setting); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	updatedSetting, err := h.settingUseCase.FindSettingByID(id)
	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(dto.ToSettingResponse(updatedSetting))
}

// DeleteSetting godoc
// @Summary Delete an setting by ID
// @Tags settings
// @Produce json
// @Param id path int true "Setting ID"
// @Success 200 {object} response.MessageResponse
// @Router /settings/{id} [delete]
func (h *HttpSettingHandler) DeleteSetting(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.settingUseCase.DeleteSetting(id); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return responses.Message(c, fiber.StatusOK, "setting deleted")
}
