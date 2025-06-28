package dto

import (
	"github.com/google/uuid"
)

type CreateItemInstanceRequest struct {
	InventoryID  uuid.UUID `gorm:"type:uuid" json:"inventory_id"`
	ItemID       uuid.UUID `gorm:"type:uuid" json:"item_id"`
	UpgradeLevel int       `json:"upgrade_level"`
}
