package dto

import "github.com/MingPV/clean-go-template/internal/entities"

func ToEquipmentSlotResponse(equipmentSlot *entities.EquipmentSlot) *EquipmentSlotResponse {
	return &EquipmentSlotResponse{
		ID:             equipmentSlot.ID,
		CharacterID:    equipmentSlot.CharacterID,
		SlotType:       equipmentSlot.SlotType,
		ItemInstanceID: equipmentSlot.ItemInstanceID,

		// ItemInstance: equipmentSlot.ItemInstance,
	}
}

func ToEquipmentSlotResponseList(equipmentSlots []*entities.EquipmentSlot) []*EquipmentSlotResponse {
	result := make([]*EquipmentSlotResponse, 0, len(equipmentSlots))
	for _, o := range equipmentSlots {
		result = append(result, ToEquipmentSlotResponse(o))
	}
	return result
}

// type EquipmentSlot struct {
// 	ID             uuid.UUID `gorm:"type:uuid;primaryKey" json:"equipment_slot_id"`
// 	CharacterID    uuid.UUID `gorm:"type:uuid" json:"character_id"`
// 	SlotType       string    `json:"slot_type"`
// 	ItemInstanceID uuid.UUID `gorm:"type:uuid" json:"item_instance_id"`

// 	ItemInstance ItemInstance `gorm:"foreignKey:ItemInstanceID;references:ID" json:"item_instance"` // this.ItemInstanceID -> ItemInstance.ID
// }
