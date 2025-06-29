package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MonsterType struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"monster_type_id"`
	Name string    `json:"name"`

	// Monsters []Monster `gorm:"foreignKey:MonsterTypeID;references:ID" json:"monsters"` // Monsters.MonsterTypeID -> this.ID
}

func (m *MonsterType) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}
