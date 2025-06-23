package rest

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/MingPV/clean-go-template/internal/rarity/dto"
	"github.com/MingPV/clean-go-template/internal/rarity/usecase"
	responses "github.com/MingPV/clean-go-template/pkg/responses"
	"github.com/gofiber/fiber/v2"
)

type HttpRarityHandler struct {
	rarityUseCase usecase.RarityUseCase
}

func NewHttpRarityHandler(useCase usecase.RarityUseCase) *HttpRarityHandler {
	return &HttpRarityHandler{rarityUseCase: useCase}
}

// CreateRarity godoc
// @Summary Create a new rarity
// @Tags rarities
// @Accept json
// @Produce json
// @Param rarity body entities.Rarity true "Rarity payload"
// @Success 201 {object} entities.Rarity
// @Router /rarities [post]
func (h *HttpRarityHandler) CreateRarity(c *fiber.Ctx) error {
	var req dto.CreateRarityRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.Error(c, fiber.StatusBadRequest, "invalid request")
	}

	rarity := &entities.Rarity{Name: req.Name, DropRate: req.DropRate}
	if err := h.rarityUseCase.CreateRarity(rarity); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToRarityResponse(rarity))
}

// FindAllRarities godoc
// @Summary Get all rarities
// @Tags rarities
// @Produce json
// @Success 200 {array} entities.Rarity
// @Router /rarities [get]
func (h *HttpRarityHandler) FindAllRarities(c *fiber.Ctx) error {
	rarities, err := h.rarityUseCase.FindAllRarities()
	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(dto.ToRarityResponseList(rarities))
}

// FindRarityByID godoc
// @Summary Get rarity by ID
// @Tags rarities
// @Produce json
// @Param id path int true "Rarity ID"
// @Success 200 {object} entities.Rarity
// @Router /rarities/{id} [get]
func (h *HttpRarityHandler) FindRarityByID(c *fiber.Ctx) error {
	id := c.Params("id")

	rarity, err := h.rarityUseCase.FindRarityByID(id)
	if err != nil {
		return responses.Error(c, fiber.StatusNotFound, err.Error())
	}

	return c.JSON(dto.ToRarityResponse(rarity))
}

// PatchRarity godoc
// @Summary Update an rarity partially
// @Tags rarities
// @Accept json
// @Produce json
// @Param id path int true "Rarity ID"
// @Param rarity body entities.Rarity true "Rarity update payload"
// @Success 200 {object} entities.Rarity
// @Router /rarities/{id} [patch]
func (h *HttpRarityHandler) PatchRarity(c *fiber.Ctx) error {
	id := c.Params("id")

	var req dto.CreateRarityRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.Error(c, fiber.StatusBadRequest, "invalid request")
	}

	rarity := &entities.Rarity{Name: req.Name}
	if err := h.rarityUseCase.PatchRarity(id, rarity); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	updatedRarity, err := h.rarityUseCase.FindRarityByID(id)
	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(dto.ToRarityResponse(updatedRarity))
}

// DeleteRarity godoc
// @Summary Delete an rarity by ID
// @Tags rarities
// @Produce json
// @Param id path int true "Rarity ID"
// @Success 200 {object} response.MessageResponse
// @Router /rarities/{id} [delete]
func (h *HttpRarityHandler) DeleteRarity(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.rarityUseCase.DeleteRarity(id); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return responses.Message(c, fiber.StatusOK, "rarity deleted")
}
