package dto

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/google/uuid"
)

type MonsterLootResponse struct {
	MonsterID   uuid.UUID `gorm:"type:uuid" json:"id"`
	ItemID      uuid.UUID `gorm:"type:uuid" json:"item_id"`
	DropRate    float64   `json:"drop_rate"`
	QuantityMin int       `json:"quantity_min"`
	QuantityMax int       `json:"quantity_max"`

	Item entities.Item `json:"item"` // this.ItemID -> Item.ID
}
