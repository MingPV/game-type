package dto

import "github.com/google/uuid"

type InventoryResponse struct {
	// ID    uint    `json:"id"`
	// Total float64 `json:"total"`
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"inventory_id"`
	MaxSlots int       `json:"max_slots"`

	ItemInstance []ItemInstance `json:"item_instance"`
}

// type Inventory struct {
// 	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"inventory_id"`
// 	MaxSlots int       `json:"max_slots"`

// 	ItemInstance []ItemInstance `gorm:"foreignKey:InventoryID;references:ID" json:"item_instance"` // ItemInstance.InventoryID -> this.ID
// }
