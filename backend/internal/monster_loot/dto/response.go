package dto

import (
	itemDTO "github.com/MingPV/clean-go-template/internal/item/dto"
	"github.com/google/uuid"
)

type MonsterLootResponse struct {
	MonsterID   uuid.UUID `gorm:"type:uuid" json:"monster_id"`
	ItemID      uuid.UUID `gorm:"type:uuid" json:"item_id"`
	QuantityMin int       `json:"quantity_min"`
	QuantityMax int       `json:"quantity_max"`

	Item itemDTO.ItemResponse `json:"item"` // this.ItemID -> Item.ID
}
