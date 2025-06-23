package rest

import (
	"github.com/MingPV/clean-go-template/internal/character/dto"
	characterUseCase "github.com/MingPV/clean-go-template/internal/character/usecase"
	"github.com/google/uuid"

	"github.com/MingPV/clean-go-template/internal/entities"
	responses "github.com/MingPV/clean-go-template/pkg/responses"
	"github.com/gofiber/fiber/v2"
)

type HttpCharacterHandler struct {
	characterUseCase characterUseCase.CharacterUseCase
}

func NewHttpCharacterHandler(character_useCase characterUseCase.CharacterUseCase) *HttpCharacterHandler {
	return &HttpCharacterHandler{characterUseCase: character_useCase}
}

// CreateCharacter godoc
// @Summary Create a new character
// @Tags characters
// @Accept json
// @Produce json
// @Param character body entities.Character true "Character payload"
// @Success 201 {object} entities.Character
// @Router /characters [post]
func (h *HttpCharacterHandler) CreateCharacter(c *fiber.Ctx) error {
	var req dto.CreateCharacterRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.Error(c, fiber.StatusBadRequest, "invalid request")
	}

	// use in base_status, character
	character_id := uuid.New()

	character := &entities.Character{
		ID:         character_id,
		UserID:     req.UserID,
		Name:       req.Name,
		Level:      req.Level,
		CurrentExp: req.CurrentExp,
		ClassID:    req.ClassID,
		StatusID:   character_id,
	}

	character_return, err := h.characterUseCase.CreateCharacter(character)

	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToCharacterResponse(character_return))
}

// FindAllCharacters godoc
// @Summary Get all characters
// @Tags characters
// @Produce json
// @Success 200 {array} entities.Character
// @Router /characters [get]
func (h *HttpCharacterHandler) FindAllCharacters(c *fiber.Ctx) error {
	characters, err := h.characterUseCase.FindAllCharacters()
	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(dto.ToCharacterResponseList(characters))
}

// FindCharacterByID godoc
// @Summary Get character by ID
// @Tags characters
// @Produce json
// @Param id path int true "Character ID"
// @Success 200 {object} entities.Character
// @Router /characters/{id} [get]
func (h *HttpCharacterHandler) FindCharacterByID(c *fiber.Ctx) error {
	id := c.Params("id")
	// characterID, err := id
	// if err != nil {
	// 	return responses.Error(c, fiber.StatusBadRequest, "invalid id")
	// }

	character, err := h.characterUseCase.FindCharacterByID(id)
	if err != nil {
		return responses.Error(c, fiber.StatusNotFound, err.Error())
	}

	return c.JSON(dto.ToCharacterResponse(character))
}

// PatchCharacter godoc
// @Summary Update an character partially
// @Tags characters
// @Accept json
// @Produce json
// @Param id path int true "Character ID"
// @Param character body entities.Character true "Character update payload"
// @Success 200 {object} entities.Character
// @Router /characters/{id} [patch]
// func (h *HttpCharacterHandler) PatchCharacter(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	characterID, err := strconv.Atoi(id)
// 	if err != nil {
// 		return responses.Error(c, fiber.StatusBadRequest, "invalid id")
// 	}

// 	var req dto.CreateCharacterRequest
// 	if err := c.BodyParser(&req); err != nil {
// 		return responses.Error(c, fiber.StatusBadRequest, "invalid request")
// 	}

// 	character := &entities.Character{Total: req.Total}
// 	if err := h.characterUseCase.PatchCharacter(characterID, character); err != nil {
// 		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
// 	}

// 	updatedCharacter, err := h.characterUseCase.FindCharacterByID(characterID)
// 	if err != nil {
// 		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
// 	}

// 	return c.JSON(dto.ToCharacterResponse(updatedCharacter))
// }

// DeleteCharacter godoc
// @Summary Delete an character by ID
// @Tags characters
// @Produce json
// @Param id path int true "Character ID"
// @Success 200 {object} response.MessageResponse
// @Router /characters/{id} [delete]
// func (h *HttpCharacterHandler) DeleteCharacter(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	characterID, err := strconv.Atoi(id)
// 	if err != nil {
// 		return responses.Error(c, fiber.StatusBadRequest, "invalid id")
// 	}

// 	if err := h.characterUseCase.DeleteCharacter(characterID); err != nil {
// 		return responses.Error(c, fiber.StatusInternalServerError, err.Error())
// 	}

// 	return responses.Message(c, fiber.StatusOK, "character deleted")
// }
