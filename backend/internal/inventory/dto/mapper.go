package dto

import "github.com/MingPV/clean-go-template/internal/entities"

func ToInventoryResponse(inventory *entities.Inventory) *InventoryResponse {
	return &InventoryResponse{
		ID:       inventory.ID,
		MaxSlots: inventory.MaxSlots,
	}
}

func ToInventoryResponseList(inventories []*entities.Inventory) []*InventoryResponse {
	result := make([]*InventoryResponse, 0, len(inventories))
	for _, o := range inventories {
		result = append(result, ToInventoryResponse(o))
	}
	return result
}

// type Inventory struct {
// 	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"inventory_id"`
// 	MaxSlots int       `json:"max_slots"`

// 	ItemInstance []ItemInstance `gorm:"foreignKey:InventoryID;references:ID" json:"item_instance"` // ItemInstance.InventoryID -> this.ID
// }
