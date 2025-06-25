package dto

import (
	"github.com/google/uuid"
)

type RarityResponse struct {
	ID       uuid.UUID `gorm:"type:uuid" json:"rarity_id"`
	Name     string    `json:"name"`
	DropRate float64   `json:"drop_rate"`
}
