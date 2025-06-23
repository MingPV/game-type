package dto

import "github.com/google/uuid"

type CreateMonsterRequest struct {
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Level         int       `json:"level"`
	HP            int       `json:"hp"`
	Attack        int       `json:"attack"`
	Defense       int       `json:"defense"`
	ExpReward     int       `json:"exp_reward"`
	GoldReward    int       `json:"gold_reward"`
	MonsterTypeID uuid.UUID `gorm:"type:uuid" json:"monster_type_id"`
}
