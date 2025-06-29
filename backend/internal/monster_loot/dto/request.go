package dto

import "github.com/google/uuid"

type CreateMonsterLootRequest struct {
	MonsterID   uuid.UUID `gorm:"type:uuid" json:"monster_id"`
	ItemID      uuid.UUID `gorm:"type:uuid" json:"item_id"`
	QuantityMin int       `json:"quantity_min"`
	QuantityMax int       `json:"quantity_max"`
}
