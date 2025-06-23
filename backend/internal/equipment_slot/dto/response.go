package dto

import (
	"github.com/google/uuid"
)

type EquipmentSlotResponse struct {
	// ID    uint    `json:""`
	// Total float64 `json:"total"`
	ID             uuid.UUID `gorm:"type:uuid" json:"equipment_slot_id"`
	CharacterID    uuid.UUID `gorm:"type:uuid" json:"character_id"`
	SlotType       string    `json:"slot_type"`
	ItemInstanceID uuid.UUID `gorm:"type:uuid" json:"item_instance_id"`

	// ItemInstance entities.ItemInstance `json:"item_instance"`
}

// type EquipmentSlot struct {
// 	ID             uuid.UUID `gorm:"type:uuid;primaryKey" json:"equipment_slot_id"`
// 	CharacterID    uuid.UUID `gorm:"type:uuid" json:"character_id"`
// 	SlotType       string    `json:"slot_type"`
// 	ItemInstanceID uuid.UUID `gorm:"type:uuid" json:"item_instance_id"`

// 	ItemInstance ItemInstance `gorm:"foreignKey:ItemInstanceID;references:ID" json:"item_instance"` // this.ItemInstanceID -> ItemInstance.ID
// }
