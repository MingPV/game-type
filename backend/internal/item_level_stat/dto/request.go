package dto

import "github.com/google/uuid"

type CreateItemLevelStatRequest struct {
	ItemID   uuid.UUID `gorm:"type:uuid" json:"item_id"`
	Level    int       ` json:"level"`
	BonusSTR int       `json:"bonus_str"`
	BonusAGI int       `json:"bonus_agi"`
	BonusINT int       `json:"bonus_int"`
	BonusDEX int       `json:"bonus_dex"`
	BonusVIT int       `json:"bonus_vit"`
	BonusLUK int       `json:"bonus_luk"`
}
