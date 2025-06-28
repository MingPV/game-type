package dto

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/google/uuid"
)

type CreateItemRequest struct {
	Name          string                        `json:"name"`
	Description   string                        `json:"description"`
	ItemTypeID    uuid.UUID                     `gorm:"type:uuid" json:"item_type_id"`
	RarityID      uuid.UUID                     `gorm:"type:uuid" json:"rarity_id"`
	RequiredLevel int                           `json:"required_level"`
	MaxStack      int                           `json:"max_stack"`
	LevelStat     map[string]entities.BonusStat `json:"level_stat"`
}
