package rest

import (
	"github.com/MingPV/clean-go-template/internal/class/dto"
	"github.com/MingPV/clean-go-template/internal/class/usecase"

	// "github.com/MingPV/clean-go-template/internal/class/usecase"
	"github.com/MingPV/clean-go-template/internal/entities"
	responses "github.com/MingPV/clean-go-template/pkg/responses"
	"github.com/gofiber/fiber/v2"
)

type HttpClassHandler struct {
	classUseCase usecase.ClassUseCase
}

func NewHttpClassHandler(useCase usecase.ClassUseCase) *HttpClassHandler {
	return &HttpClassHandler{classUseCase: useCase}
}

// CreateClass godoc
// @Summary Create a new class
// @Tags classes
// @Accept json
// @Produce json
// @Param class body entities.Class true "Class payload"
// @Success 201 {object} entities.Class
// @Router /classes [post]
func (h *HttpClassHandler) CreateClass(c *fiber.Ctx) error {
	var req dto.CreateClassRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.Error(c, fiber.StatusBadRequest, "invalid request")
	}

	class := &entities.Class{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
	}

	if err := h.classUseCase.CreateClass(class); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToClassResponse(class))
}

// FindAllClasses godoc
// @Summary Get all classes
// @Tags classes
// @Produce json
// @Success 200 {array} entities.Class
// @Router /classes [get]
func (h *HttpClassHandler) FindAllClasses(c *fiber.Ctx) error {
	classes, err := h.classUseCase.FindAllClasses()
	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(dto.ToClassResponseList(classes))
}

// FindClassByID godoc
// @Summary Get class by ID
// @Tags classes
// @Produce json
// @Param id path int true "Class ID"
// @Success 200 {object} entities.Class
// @Router /classes/{id} [get]
func (h *HttpClassHandler) FindClassByID(c *fiber.Ctx) error {
	id := c.Params("id")

	class, err := h.classUseCase.FindClassByID(id)
	if err != nil {
		return responses.Error(c, fiber.StatusNotFound, err.Error())
	}

	return c.JSON(dto.ToClassResponse(class))
}

// PatchClass godoc
// @Summary Update an class partially
// @Tags classes
// @Accept json
// @Produce json
// @Param id path int true "Class ID"
// @Param class body entities.Class true "Class update payload"
// @Success 200 {object} entities.Class
// @Router /classes/{id} [patch]
// func (h *HttpClassHandler) PatchClass(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	classID, err := strconv.Atoi(id)
// 	if err != nil {
// 		return responses.Error(c, fiber.StatusBadRequest, "invalid id")
// 	}

// 	var req dto.CreateClassRequest
// 	if err := c.BodyParser(&req); err != nil {
// 		return responses.Error(c, fiber.StatusBadRequest, "invalid request")
// 	}

// 	class := &entities.Class{Total: req.Total}
// 	if err := h.classUseCase.PatchClass(classID, class); err != nil {
// 		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
// 	}

// 	updatedClass, err := h.classUseCase.FindClassByID(classID)
// 	if err != nil {
// 		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
// 	}

// 	return c.JSON(dto.ToClassResponse(updatedClass))
// }

// DeleteClass godoc
// @Summary Delete an class by ID
// @Tags classes
// @Produce json
// @Param id path int true "Class ID"
// @Success 200 {object} response.MessageResponse
// @Router /classes/{id} [delete]
// func (h *HttpStatusHandler) DeleteClass(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	classID, err := strconv.Atoi(id)
// 	if err != nil {
// 		return responses.Error(c, fiber.StatusBadRequest, "invalid id")
// 	}

// 	if err := h.classUseCase.DeleteClass(classID); err != nil {
// 		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
// 	}

// 	return responses.Message(c, fiber.StatusOK, "class deleted")
// }
