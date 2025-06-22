package dto

import "github.com/MingPV/clean-go-template/internal/entities"

func ToItemResponse(item *entities.Item) *ItemResponse {
	return &ItemResponse{
		ID:            item.ID,
		Name:          item.Name,
		Description:   item.Description,
		ItemTypeID:    item.ItemTypeID,
		RarityID:      item.RarityID,
		RequiredLevel: item.RequiredLevel,
		MaxStack:      item.MaxStack,
		// ItemType:  ToItemTypeResponse(item.ItemType),
		// Rarity:    ToRarityResponse(item.Rarity),
		// ItemStats: ToItemLevelStatResponseList(item.ItemStats),
		ItemType:  item.ItemType,
		Rarity:    item.Rarity,
		ItemStats: item.ItemStats,
	}
}

func ToItemResponseList(items []*entities.Item) []*ItemResponse {
	result := make([]*ItemResponse, 0, len(items))
	for _, o := range items {
		result = append(result, ToItemResponse(o))
	}
	return result
}

// type Item struct {
// 	ID            uuid.UUID `gorm:"type:uuid;primaryKey" json:"item_id"`
// 	Name          string    `json:"name"`
// 	Description   string    `json:"description"`
// 	ItemTypeID    uuid.UUID `gorm:"type:uuid" json:"item_type_id"`
// 	RarityID      uuid.UUID `gorm:"type:uuid" json:"rarity_id"`
// 	RequiredLevel int       `json:"required_level"`
// 	MaxStack      int       `json:"max_stack"`

// 	ItemType  ItemType        `gorm:"foreignKey:ItemTypeID;references:ID" json:"item_type"`   // this.ItemTypeID -> ItemType.ID
// 	Rarity    Rarity          `gorm:"foreignKey:RarityID;references:ID" json:"rarity"`        // this.RarityID -> Rarity.ID
// 	ItemStats []ItemLevelStat `gorm:"foreignKey:ItemID;references:ID" json:"item_level_stat"` // ItemLevelStat.ItemID -> this.ID
// }
