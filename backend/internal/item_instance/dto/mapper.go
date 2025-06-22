package dto

import "github.com/MingPV/clean-go-template/internal/entities"

func ToItemInstanceResponse(itemInstance *entities.ItemInstance) *ItemInstanceResponse {
	return &ItemInstanceResponse{
		ID:               itemInstance.ID,
		InventoryID:      itemInstance.InventoryID,
		ItemID:           itemInstance.ItemID,
		UpgradeLevel:     itemInstance.UpgradeLevel,
		OwnerCharacterID: itemInstance.OwnerCharacterID,
		CreatedAt:        itemInstance.CreatedAt,
		Item:             itemInstance.Item, // this.ItemID -> Item.ID
	}
}

func ToItemInstanceResponseList(itemInstances []*entities.ItemInstance) []*ItemInstanceResponse {
	result := make([]*ItemInstanceResponse, 0, len(itemInstances))
	for _, o := range itemInstances {
		result = append(result, ToItemInstanceResponse(o))
	}
	return result
}

// 	ID               uuid.UUID `gorm:"type:uuid;primaryKey" json:"item_instance_id"`
// 	InventoryID      uuid.UUID `gorm:"type:uuid" json:"inventory_id"`
// 	ItemID           uuid.UUID `gorm:"type:uuid" json:"item_id"`
// 	UpgradeLevel     int       `json:"upgrade_level"`
// 	OwnerCharacterID uuid.UUID `gorm:"type:uuid" json:"owner_character_id"`
// 	CreatedAt        time.Time `json:"created_at"`

// 	Item Item `gorm:"foreignKey:ItemID;references:ID" json:"item"` // this.ItemID -> Item.ID
