package dto

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	itemInstanceDTO "github.com/MingPV/clean-go-template/internal/item_instance/dto"
)

func ToInventoryResponse(inventory *entities.Inventory) *InventoryResponse {

	itemInstances := make([]itemInstanceDTO.ItemInstanceResponse, 0, len(inventory.ItemInstances))
	for _, ii := range inventory.ItemInstances {
		itemInstances = append(itemInstances, itemInstanceDTO.ToItemInstanceResponse2(ii))
	}

	return &InventoryResponse{
		ID:            inventory.ID,
		MaxSlots:      inventory.MaxSlots,
		ItemInstances: itemInstances,
	}
}

func ToInventoryResponseList(inventories []*entities.Inventory) []*InventoryResponse {
	result := make([]*InventoryResponse, 0, len(inventories))
	for _, o := range inventories {
		result = append(result, ToInventoryResponse(o))
	}
	return result
}
