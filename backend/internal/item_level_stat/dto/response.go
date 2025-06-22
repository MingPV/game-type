package dto

import "github.com/google/uuid"

type ItemLevelStatResponse struct {
	ItemID   uuid.UUID `gorm:"type:uuid" json:"item_id"`
	Level    int       `json:"level"`
	BonusSTR int       `json:"bonus_str"`
	BonusAGI int       `json:"bonus_agi"`
	BonusINT int       `json:"bonus_int"`
	BonusDEX int       `json:"bonus_dex"`
	BonusVIT int       `json:"bonus_vit"`
	BonusLUK int       `json:"bonus_luk"`
}

// type ItemLevelStat struct {
// 	ItemID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"item_id"`
// 	Level    int       `gorm:"primaryKey" json:"level"`
// 	BonusSTR int       `json:"bonus_str"`
// 	BonusAGI int       `json:"bonus_agi"`
// 	BonusINT int       `json:"bonus_int"`
// 	BonusDEX int       `json:"bonus_dex"`
// 	BonusVIT int       `json:"bonus_vit"`
// 	BonusLUK int       `json:"bonus_luk"`

// 	Item Item `gorm:"foreignKey:ItemID;references:ID"` // this.ItemID -> Item.ID
// }
