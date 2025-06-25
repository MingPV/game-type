package dto

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/MingPV/clean-go-template/internal/item_level_stat/dto"
)

func ToItemResponse(item *entities.Item) *ItemResponse {

	return &ItemResponse{
		ID:            item.ID,
		Name:          item.Name,
		Description:   item.Description,
		ItemTypeID:    item.ItemTypeID,
		RarityID:      item.RarityID,
		RequiredLevel: item.RequiredLevel,
		MaxStack:      item.MaxStack,

		ItemType:  item.ItemType,
		Rarity:    item.Rarity,
		ItemStats: *dto.ToItemLevelStatResponse(&item.ItemStats),
	}
}

func ToItemResponseList(items []*entities.Item) []*ItemResponse {
	result := make([]*ItemResponse, 0, len(items))
	for _, o := range items {
		result = append(result, ToItemResponse(o))
	}
	return result
}
