package dto

import (
	"time"

	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/google/uuid"
)

type ItemInstanceResponse struct {
	ID               uuid.UUID `gorm:"type:uuid" json:"item_instance_id"`
	InventoryID      uuid.UUID `gorm:"type:uuid" json:"inventory_id"`
	ItemID           uuid.UUID `gorm:"type:uuid" json:"item_id"`
	UpgradeLevel     int       `json:"upgrade_level"`
	OwnerCharacterID uuid.UUID `gorm:"type:uuid" json:"owner_character_id"`
	CreatedAt        time.Time `json:"created_at"`

	Item entities.Item ` json:"item"` // this.ItemID -> Item.ID
}

// type ItemInstance struct {
// 	ID               uuid.UUID `gorm:"type:uuid;primaryKey" json:"item_instance_id"`
// 	InventoryID      uuid.UUID `gorm:"type:uuid" json:"inventory_id"`
// 	ItemID           uuid.UUID `gorm:"type:uuid" json:"item_id"`
// 	UpgradeLevel     int       `json:"upgrade_level"`
// 	OwnerCharacterID uuid.UUID `gorm:"type:uuid" json:"owner_character_id"`
// 	CreatedAt        time.Time `json:"created_at"`

// 	Item Item `gorm:"foreignKey:ItemID;references:ID" json:"item"` // this.ItemID -> Item.ID
// }
