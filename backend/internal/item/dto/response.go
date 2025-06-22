package dto

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/google/uuid"
)

type ItemResponse struct {
	// ID    uint    `json:"id"`
	// Total float64 `json:"total"`
	ID            uuid.UUID `gorm:"type:uuid" json:"item_id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	ItemTypeID    uuid.UUID `gorm:"type:uuid" json:"item_type_id"`
	RarityID      uuid.UUID `gorm:"type:uuid" json:"rarity_id"`
	RequiredLevel int       `json:"required_level"`
	MaxStack      int       `json:"max_stack"`

	ItemType  entities.ItemType        `json:"item_type"`       // this.ItemTypeID -> ItemType.ID
	Rarity    entities.Rarity          `json:"rarity"`          // this.RarityID -> Rarity.ID
	ItemStats []entities.ItemLevelStat `json:"item_level_stat"` // ItemLevelStat.ItemID -> this.ID
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
