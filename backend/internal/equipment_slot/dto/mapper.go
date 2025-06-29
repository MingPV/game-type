package dto

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	itemInstanceDTO "github.com/MingPV/clean-go-template/internal/item_instance/dto"
)

func ToEquipmentSlotResponse(equipmentSlot *entities.EquipmentSlot) *EquipmentSlotResponse {
	return &EquipmentSlotResponse{
		ID:             equipmentSlot.ID,
		CharacterID:    equipmentSlot.CharacterID,
		SlotType:       equipmentSlot.SlotType,
		ItemInstanceID: equipmentSlot.ItemInstanceID,

		ItemInstance: *itemInstanceDTO.ToItemInstanceResponse(&equipmentSlot.ItemInstance),
	}
}

func ToEquipmentSlotResponse2(equipmentSlot entities.EquipmentSlot) EquipmentSlotResponse {
	return EquipmentSlotResponse{
		ID:             equipmentSlot.ID,
		CharacterID:    equipmentSlot.CharacterID,
		SlotType:       equipmentSlot.SlotType,
		ItemInstanceID: equipmentSlot.ItemInstanceID,

		ItemInstance: *itemInstanceDTO.ToItemInstanceResponse(&equipmentSlot.ItemInstance),
	}
}

func ToEquipmentSlotResponseList(equipmentSlots []*entities.EquipmentSlot) []*EquipmentSlotResponse {
	result := make([]*EquipmentSlotResponse, 0, len(equipmentSlots))
	for _, o := range equipmentSlots {
		result = append(result, ToEquipmentSlotResponse(o))
	}
	return result
}
