package dto

import "github.com/google/uuid"

type CreateEquipmentSlotRequest struct {
	CharacterID    uuid.UUID `gorm:"type:uuid" json:"character_id"`
	SlotType       string    `json:"slot_type"`
	ItemInstanceID uuid.UUID `gorm:"type:uuid" json:"item_instance_id"`
}
