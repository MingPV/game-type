package dto

import (
	"github.com/google/uuid"
)

type StatusResponse struct {
	CharacterID uuid.UUID `gorm:"type:uuid" json:"character_id"`
	StatusPoint int       `json:"status_point"`
	Attack      int       `json:"attack"`
	Defense     int       `json:"defense"`
	HP          int       `json:"hp"`
	MP          int       `json:"mp"`
	Critical    int       `json:"critical"`
}
