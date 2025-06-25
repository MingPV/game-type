package dto

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/google/uuid"
)

type CreateItemLevelStatRequest struct {
	ItemID    uuid.UUID                     `gorm:"type:uuid" json:"item_id"`
	BonusStat map[string]entities.BonusStat `json:"bonus_stat"`
}

type PatchItemLevelStatRequest struct {
	BonusStat map[string]entities.BonusStat `json:"bonus_stat"`
}
