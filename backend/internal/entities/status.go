package entities

import (
	"github.com/google/uuid"
)

type Status struct {
	CharacterID uuid.UUID `gorm:"type:uuid;primaryKey" json:"character_id"`
	StatusPoint int       `json:"status_point"`
	STR         int       `json:"str"`
	AGI         int       `json:"agi"`
	INT         int       `json:"int"`
	DEX         int       `json:"dex"`
	VIT         int       `json:"vit"`
	LUK         int       `json:"luk"`
}

// Tested
