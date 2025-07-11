package dto

import (
	"time"

	"github.com/MingPV/clean-go-template/internal/entities"
	equipmentSlotDTO "github.com/MingPV/clean-go-template/internal/equipment_slot/dto"
	statusDTO "github.com/MingPV/clean-go-template/internal/status/dto"

	// inventoryDTO "github.com/MingPV/clean-go-template/internal/inventory/dto"
	"github.com/google/uuid"
)

type CharacterResponse struct {
	ID          uuid.UUID `gorm:"type:uuid" json:"character_id"`
	UserID      uuid.UUID `gorm:"type:uuid" json:"user_id"`
	Name        string    `json:"name"`
	Level       int       `json:"level"`
	CurrentExp  int       `json:"current_exp"`
	ClassID     uuid.UUID `gorm:"type:uuid" json:"class_id"`
	CreatedAt   time.Time `json:"created_at"`
	InventoryID uuid.UUID `json:"inventory_id"`

	Class          entities.Class                           `json:"class"`
	Status         statusDTO.StatusResponse                 `json:"status"`
	EquipmentSlots []equipmentSlotDTO.EquipmentSlotResponse `json:"equipment_slots"`
	// Inventory      inventoryDTO.InventoryResponse           `json:"inventory"`
}

// type Character struct {
// 	ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"character_id"`
// 	UserID     uuid.UUID `gorm:"type:uuid" json:"user_id"`
// 	Name       string    `json:"name"`
// 	Level      int       `json:"level"`
// 	CurrentExp int       `json:"current_exp"`
// 	ClassID    uuid.UUID `gorm:"type:uuid" json:"class_id"`
// 	StatusID   uuid.UUID `gorm:"type:uuid" json:"status_id"`
// 	CreatedAt  time.Time `json:"created_at"`

// 	User   User            `gorm:"foreignKey:UserID;references:ID"`
// 	Class  Class           `gorm:"foreignKey:ClassID;references:ID" json:"class"`
// 	Status Status          `gorm:"foreignKey:StatusID;references:CharacterID" json:"status"`
// 	Slots  []EquipmentSlot `gorm:"foreignKey:CharacterID" json:"equipment_slots"`
// 	Items  []ItemInstance  `gorm:"foreignKey:OwnerCharacterID" json:"items"`
// }
