package dto

import (
	"github.com/google/uuid"
)

type CreateItemInstanceRequest struct {
	InventoryID      uuid.UUID `gorm:"type:uuid" json:"inventory_id"`
	ItemID           uuid.UUID `gorm:"type:uuid" json:"item_id"`
	UpgradeLevel     int       `json:"upgrade_level"`
	OwnerCharacterID uuid.UUID `gorm:"type:uuid" json:"owner_character_id"`
}

// 	ID               uuid.UUID `gorm:"type:uuid;primaryKey" json:"item_instance_id"`
// 	InventoryID      uuid.UUID `gorm:"type:uuid" json:"inventory_id"`
// 	ItemID           uuid.UUID `gorm:"type:uuid" json:"item_id"`
// 	UpgradeLevel     int       `json:"upgrade_level"`
// 	OwnerCharacterID uuid.UUID `gorm:"type:uuid" json:"owner_character_id"`
// 	CreatedAt        time.Time `json:"created_at"`

// 	Item Item `gorm:"foreignKey:ItemID;references:ID" json:"item"` // this.ItemID -> Item.ID
