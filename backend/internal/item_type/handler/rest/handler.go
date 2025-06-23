package rest

import (
	"strconv"

	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/MingPV/clean-go-template/internal/item_type/dto"
	"github.com/MingPV/clean-go-template/internal/item_type/usecase"
	responses "github.com/MingPV/clean-go-template/pkg/responses"
	"github.com/gofiber/fiber/v2"
)

type HttpItemTypeHandler struct {
	itemTypeUseCase usecase.ItemTypeUseCase
}

func NewHttpItemTypeHandler(useCase usecase.ItemTypeUseCase) *HttpItemTypeHandler {
	return &HttpItemTypeHandler{itemTypeUseCase: useCase}
}

// CreateItemType godoc
// @Summary Create a new itemType
// @Tags itemTypes
// @Accept json
// @Produce json
// @Param itemType body entities.ItemType true "ItemType payload"
// @Success 201 {object} entities.ItemType
// @Router /itemTypes [post]
func (h *HttpItemTypeHandler) CreateItemType(c *fiber.Ctx) error {
	var req dto.CreateItemTypeRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.Error(c, fiber.StatusBadRequest, "invalid request")
	}

	itemType := &entities.ItemType{Name: req.Name}
	if err := h.itemTypeUseCase.CreateItemType(itemType); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToItemTypeResponse(itemType))
}

// FindAllItemTypes godoc
// @Summary Get all itemTypes
// @Tags itemTypes
// @Produce json
// @Success 200 {array} entities.ItemType
// @Router /itemTypes [get]
func (h *HttpItemTypeHandler) FindAllItemTypes(c *fiber.Ctx) error {
	itemTypes, err := h.itemTypeUseCase.FindAllItemTypes()
	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(dto.ToItemTypeResponseList(itemTypes))
}

// FindItemTypeByID godoc
// @Summary Get itemType by ID
// @Tags itemTypes
// @Produce json
// @Param id path int true "ItemType ID"
// @Success 200 {object} entities.ItemType
// @Router /itemTypes/{id} [get]
func (h *HttpItemTypeHandler) FindItemTypeByID(c *fiber.Ctx) error {
	id := c.Params("id")

	itemType, err := h.itemTypeUseCase.FindItemTypeByID(id)
	if err != nil {
		return responses.Error(c, fiber.StatusNotFound, err.Error())
	}

	return c.JSON(dto.ToItemTypeResponse(itemType))
}

// PatchItemType godoc
// @Summary Update an itemType partially
// @Tags itemTypes
// @Accept json
// @Produce json
// @Param id path int true "ItemType ID"
// @Param itemType body entities.ItemType true "ItemType update payload"
// @Success 200 {object} entities.ItemType
// @Router /itemTypes/{id} [patch]
// func (h *HttpItemTypeHandler) PatchItemType(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	itemTypeID, err := strconv.Atoi(id)
// 	if err != nil {
// 		return responses.Error(c, fiber.StatusBadRequest, "invalid id")
// 	}

// 	var req dto.CreateItemTypeRequest
// 	if err := c.BodyParser(&req); err != nil {
// 		return responses.Error(c, fiber.StatusBadRequest, "invalid request")
// 	}

// 	itemType := &entities.ItemType{Total: req.Total}
// 	if err := h.itemTypeUseCase.PatchItemType(itemTypeID, itemType); err != nil {
// 		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
// 	}

// 	updatedItemType, err := h.itemTypeUseCase.FindItemTypeByID(itemTypeID)
// 	if err != nil {
// 		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
// 	}

// 	return c.JSON(dto.ToItemTypeResponse(updatedItemType))
// }

// DeleteItemType godoc
// @Summary Delete an itemType by ID
// @Tags itemTypes
// @Produce json
// @Param id path int true "ItemType ID"
// @Success 200 {object} response.MessageResponse
// @Router /itemTypes/{id} [delete]
func (h *HttpItemTypeHandler) DeleteItemType(c *fiber.Ctx) error {
	id := c.Params("id")
	itemTypeID, err := strconv.Atoi(id)
	if err != nil {
		return responses.Error(c, fiber.StatusBadRequest, "invalid id")
	}

	if err := h.itemTypeUseCase.DeleteItemType(itemTypeID); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return responses.Message(c, fiber.StatusOK, "itemType deleted")
}
