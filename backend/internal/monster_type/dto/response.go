package dto

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/google/uuid"
)

type MonsterTypeResponse struct {
	ID   uuid.UUID `gorm:"type:uuid" json:"monster_type_id"`
	Name string    `json:"name"`

	Monsters []entities.Monster `json:"monsters"` // Monsters.MonsterTypeID -> this.ID
}
