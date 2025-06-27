package entities

import "github.com/google/uuid"

type ItemLevelStat struct {
	ItemID    uuid.UUID `gorm:"type:uuid;primaryKey" json:"item_id"`
	BonusStat string    `json:"bonus_stat"`
}

type BonusStat struct {
	BonusSTR int `json:"bonus_str"`
	BonusAGI int `json:"bonus_agi"`
	BonusINT int `json:"bonus_int"`
	BonusDEX int `json:"bonus_dex"`
	BonusVIT int `json:"bonus_vit"`
	BonusLUK int `json:"bonus_luk"`
}

// Tested
