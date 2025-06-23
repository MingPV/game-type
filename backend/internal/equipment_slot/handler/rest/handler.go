package rest

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/MingPV/clean-go-template/internal/equipment_slot/dto"
	"github.com/MingPV/clean-go-template/internal/equipment_slot/usecase"
	responses "github.com/MingPV/clean-go-template/pkg/responses"
	"github.com/gofiber/fiber/v2"
)

type HttpEquipmentSlotHandler struct {
	equipmentSlotUseCase usecase.EquipmentSlotUseCase
}

func NewHttpEquipmentSlotHandler(useCase usecase.EquipmentSlotUseCase) *HttpEquipmentSlotHandler {
	return &HttpEquipmentSlotHandler{equipmentSlotUseCase: useCase}
}

// CreateEquipmentSlot godoc
// @Summary Create a new equipmentSlot
// @Tags equipmentSlots
// @Accept json
// @Produce json
// @Param equipmentSlot body entities.EquipmentSlot true "EquipmentSlot payload"
// @Success 201 {object} entities.EquipmentSlot
// @Router /equipmentSlots [post]
func (h *HttpEquipmentSlotHandler) CreateEquipmentSlot(c *fiber.Ctx) error {
	var req dto.CreateEquipmentSlotRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.Error(c, fiber.StatusBadRequest, "invalid request")
	}

	// equipmentSlot := &entities.EquipmentSlot{Total: req.Total}

	equipmentSlot := &entities.EquipmentSlot{
		CharacterID:    req.CharacterID,
		SlotType:       req.SlotType,
		ItemInstanceID: req.ItemInstanceID,
	}
	if err := h.equipmentSlotUseCase.CreateEquipmentSlot(equipmentSlot); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToEquipmentSlotResponse(equipmentSlot))
}

// FindAllEquipmentSlots godoc
// @Summary Get all equipmentSlots
// @Tags equipmentSlots
// @Produce json
// @Success 200 {array} entities.EquipmentSlot
// @Router /equipmentSlots [get]
func (h *HttpEquipmentSlotHandler) FindAllEquipmentSlots(c *fiber.Ctx) error {
	equipmentSlots, err := h.equipmentSlotUseCase.FindAllEquipmentSlots()
	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(dto.ToEquipmentSlotResponseList(equipmentSlots))
}

// FindEquipmentSlotByID godoc
// @Summary Get equipmentSlot by ID
// @Tags equipmentSlots
// @Produce json
// @Param id path int true "EquipmentSlot ID"
// @Success 200 {object} entities.EquipmentSlot
// @Router /equipmentSlots/{id} [get]
func (h *HttpEquipmentSlotHandler) FindEquipmentSlotByID(c *fiber.Ctx) error {
	id := c.Params("id")

	equipmentSlot, err := h.equipmentSlotUseCase.FindEquipmentSlotByID(id)
	if err != nil {
		return responses.Error(c, fiber.StatusNotFound, err.Error())
	}

	return c.JSON(dto.ToEquipmentSlotResponse(equipmentSlot))
}

// PatchEquipmentSlot godoc
// @Summary Update an equipmentSlot partially
// @Tags equipmentSlots
// @Accept json
// @Produce json
// @Param id path int true "EquipmentSlot ID"
// @Param equipmentSlot body entities.EquipmentSlot true "EquipmentSlot update payload"
// @Success 200 {object} entities.EquipmentSlot
// @Router /equipmentSlots/{id} [patch]
func (h *HttpEquipmentSlotHandler) PatchEquipmentSlot(c *fiber.Ctx) error {
	id := c.Params("id")

	var req dto.CreateEquipmentSlotRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.Error(c, fiber.StatusBadRequest, "invalid request")
	}

	equipmentSlot := &entities.EquipmentSlot{SlotType: req.SlotType}
	if err := h.equipmentSlotUseCase.PatchEquipmentSlot(id, equipmentSlot); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	updatedEquipmentSlot, err := h.equipmentSlotUseCase.FindEquipmentSlotByID(id)
	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(dto.ToEquipmentSlotResponse(updatedEquipmentSlot))
}

// DeleteEquipmentSlot godoc
// @Summary Delete an equipmentSlot by ID
// @Tags equipmentSlots
// @Produce json
// @Param id path int true "EquipmentSlot ID"
// @Success 200 {object} response.MessageResponse
// @Router /equipmentSlots/{id} [delete]
func (h *HttpEquipmentSlotHandler) DeleteEquipmentSlot(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.equipmentSlotUseCase.DeleteEquipmentSlot(id); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return responses.Message(c, fiber.StatusOK, "equipmentSlot deleted")
}
