package entities

import (
	"github.com/google/uuid"
)

type MonsterLoot struct {
	MonsterID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	ItemID      uuid.UUID `gorm:"type:uuid;primaryKey" json:"item_id"`
	DropRate    float64   `json:"drop_rate"`
	QuantityMin int       `json:"quantity_min"`
	QuantityMax int       `json:"quantity_max"`

	Item Item `gorm:"foreignKey:ItemID;references:ID" json:"item"` // this.ItemID -> Item.ID
}
