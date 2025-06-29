package dto

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	itemDTO "github.com/MingPV/clean-go-template/internal/item/dto"
)

func ToItemInstanceResponse(itemInstance *entities.ItemInstance) *ItemInstanceResponse {
	return &ItemInstanceResponse{
		ID:           itemInstance.ID,
		InventoryID:  itemInstance.InventoryID,
		ItemID:       itemInstance.ItemID,
		UpgradeLevel: itemInstance.UpgradeLevel,
		CreatedAt:    itemInstance.CreatedAt,
		Item:         *itemDTO.ToItemResponse(&itemInstance.Item), // this.ItemID -> Item.ID
	}
}

func ToItemInstanceResponse2(itemInstance entities.ItemInstance) ItemInstanceResponse {
	return ItemInstanceResponse{
		ID:           itemInstance.ID,
		InventoryID:  itemInstance.InventoryID,
		ItemID:       itemInstance.ItemID,
		UpgradeLevel: itemInstance.UpgradeLevel,
		CreatedAt:    itemInstance.CreatedAt,
		Item:         *itemDTO.ToItemResponse(&itemInstance.Item), // this.ItemID -> Item.ID
	}
}

func ToItemInstanceResponseList(itemInstances []*entities.ItemInstance) []*ItemInstanceResponse {
	result := make([]*ItemInstanceResponse, 0, len(itemInstances))
	for _, o := range itemInstances {
		result = append(result, ToItemInstanceResponse(o))
	}
	return result
}
