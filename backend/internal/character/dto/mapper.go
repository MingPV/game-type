package dto

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	equipmentSlotDTO "github.com/MingPV/clean-go-template/internal/equipment_slot/dto"
	// inventoryDTO "github.com/MingPV/clean-go-template/internal/inventory/dto"
)

func ToCharacterResponse(character *entities.Character) *CharacterResponse {

	equipment_slots := make([]equipmentSlotDTO.EquipmentSlotResponse, 0, len(character.EquipmentSlots))
	for _, ii := range character.EquipmentSlots {
		equipment_slots = append(equipment_slots, equipmentSlotDTO.ToEquipmentSlotResponse2(ii))
	}

	return &CharacterResponse{
		ID:          character.ID,
		UserID:      character.UserID,
		Name:        character.Name,
		Level:       character.Level,
		CurrentExp:  character.CurrentExp,
		ClassID:     character.ClassID,
		CreatedAt:   character.CreatedAt,
		InventoryID: character.InventoryID,

		Class:          character.Class,
		Status:         character.Status,
		EquipmentSlots: equipment_slots,
		// Inventory:      *inventoryDTO.ToInventoryResponse(&character.Inventory),
	}
}

func ToCharacterResponseList(characters []*entities.Character) []*CharacterResponse {
	result := make([]*CharacterResponse, 0, len(characters))
	for _, c := range characters {
		result = append(result, ToCharacterResponse(c))
	}
	return result
}
