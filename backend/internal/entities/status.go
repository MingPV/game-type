package entities

import (
	"github.com/google/uuid"
)

type Status struct {
	CharacterID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"character_id"`
	StatusPoint   int       `json:"status_point"`
	AttackLevel   int       `json:"attack_level"`
	DefenseLevel  int       `json:"defense_level"`
	HPLevel       int       `json:"hp_level"`
	MPLevel       int       `json:"mp_level"`
	CriticalLevel int       `json:"critical_level"`
}

// Tested
