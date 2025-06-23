package dto

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/google/uuid"
)

type MonsterResponse struct {
	ID            uuid.UUID `gorm:"type:uuid" json:"monster_id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Level         int       `json:"level"`
	HP            int       `json:"hp"`
	Attack        int       `json:"attack"`
	Defense       int       `json:"defense"`
	ExpReward     int       `json:"exp_reward"`
	GoldReward    int       `json:"gold_reward"`
	MonsterTypeID uuid.UUID `gorm:"type:uuid" json:"monster_type_id"`

	MonsterType  entities.MonsterType   `json:"monster_type"`  // this.MonsterTypeID -> MonsterType.ID
	MonsterLoots []entities.MonsterLoot `json:"monster_loots"` // many to many Monster, MonsterLoot
}
