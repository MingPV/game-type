package dto

import (
	"time"

	itemDTO "github.com/MingPV/clean-go-template/internal/item/dto"
	"github.com/google/uuid"
)

type ItemInstanceResponse struct {
	ID           uuid.UUID `gorm:"type:uuid" json:"item_instance_id"`
	InventoryID  uuid.UUID `gorm:"type:uuid" json:"inventory_id"`
	ItemID       uuid.UUID `gorm:"type:uuid" json:"item_id"`
	UpgradeLevel int       `json:"upgrade_level"`
	CreatedAt    time.Time `json:"created_at"`

	Item itemDTO.ItemResponse ` json:"item"` // this.ItemID -> Item.ID
}
