package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EquipmentSlot struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey" json:"equipment_slot_id"`
	CharacterID    uuid.UUID `gorm:"type:uuid" json:"character_id"`
	SlotType       string    `json:"slot_type"`
	ItemInstanceID uuid.UUID `gorm:"type:uuid" json:"item_instance_id"`

	ItemInstance ItemInstance `gorm:"foreignKey:ItemInstanceID;references:ID" json:"item_instance"` // this.ItemInstanceID -> ItemInstance.ID
}

func (e *EquipmentSlot) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.New()
	return
}
