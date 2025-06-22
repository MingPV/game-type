package dto

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/google/uuid"
)

type RarityResponse struct {
	ID       uuid.UUID `gorm:"type:uuid" json:"rarity_id"`
	Name     string    `json:"name"`
	DropRate float64   `json:"drop_rate"`

	Items []entities.Item `gorm:"foreignKey:RarityID;references:ID" json:"items"` // Item.RarityID -> this.ID
}

// type Rarity struct {
// 	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"rarity_id"`
// 	Name     string    `json:"name"`
// 	DropRate float64   `json:"drop_rate"`

// 	Items []Item `gorm:"foreignKey:RarityID;references:ID" json:"items"` // Item.RarityID -> this.ID
// }
