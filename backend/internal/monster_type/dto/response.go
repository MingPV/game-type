package dto

import (
	"github.com/google/uuid"
)

type MonsterTypeResponse struct {
	ID   uuid.UUID `gorm:"type:uuid" json:"monster_type_id"`
	Name string    `json:"name"`
}
