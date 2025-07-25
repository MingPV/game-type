package entities

import (
	"github.com/google/uuid"
)

type Item struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey" json:"item_id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	ItemTypeID    uuid.UUID `gorm:"type:uuid" json:"item_type_id"`
	RarityID      uuid.UUID `gorm:"type:uuid" json:"rarity_id"`
	RequiredLevel int       `json:"required_level"`
	MaxStack      int       `json:"max_stack"`

	ItemType  ItemType      `gorm:"foreignKey:ItemTypeID;references:ID" json:"item_type"`   // this.ItemTypeID -> ItemType.ID
	Rarity    Rarity        `gorm:"foreignKey:RarityID;references:ID" json:"rarity"`        // this.RarityID -> Rarity.ID
	ItemStats ItemLevelStat `gorm:"foreignKey:ID;references:ItemID" json:"item_level_stat"` // this.ID -> ItemLevelStat.ItemID
}

// func (i *Item) BeforeCreate(tx *gorm.DB) (err error) {
// 	i.ID = uuid.New()
// 	return
// }

// Tested
