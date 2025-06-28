package dto

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/google/uuid"
)

type InventoryResponse struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"inventory_id"`
	MaxSlots int       `json:"max_slots"`

	ItemInstance []entities.ItemInstance `json:"item_instances"`
}
