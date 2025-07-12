package entities

import (
	"github.com/google/uuid"
)

type Status struct {
	CharacterID uuid.UUID `gorm:"type:uuid;primaryKey" json:"character_id"`
	StatusPoint int       `json:"status_point"`
	Attack      int       `json:"attack"`
	Defense     int       `json:"defense"`
	HP          int       `json:"hp"`
	MP          int       `json:"mp"`
	Critical    int       `json:"critical"`
}

// Tested
