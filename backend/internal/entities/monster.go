package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Monster struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey" json:"monster_id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Level         int       `json:"level"`
	HP            int       `json:"hp"`
	Attack        int       `json:"attack"`
	Defense       int       `json:"defense"`
	ExpReward     int       `json:"exp_reward"`
	GoldReward    int       `json:"gold_reward"`
	MonsterTypeID uuid.UUID `gorm:"type:uuid" json:"monster_type_id"`

	MonsterType  MonsterType   `gorm:"foreignKey:MonsterTypeID;references:ID" json:"monster_type"` // this.MonsterTypeID -> MonsterType.ID
	MonsterLoots []MonsterLoot `gorm:"foreignKey:MonsterID" json:"monster_loots"`                  // MonsterLoot.MonsterID -> this.ID
}

func (m *Monster) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}
