package dto

import (
	itemInstanceDTO "github.com/MingPV/clean-go-template/internal/item_instance/dto"
	"github.com/google/uuid"
)

type InventoryResponse struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"inventory_id"`
	MaxSlots int       `json:"max_slots"`

	ItemInstances []itemInstanceDTO.ItemInstanceResponse `json:"item_instances"`
}
