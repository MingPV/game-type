package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateCharacterRequest struct {
	UserID     uuid.UUID `gorm:"type:uuid" json:"user_id"`
	Name       string    `json:"name"`
	Level      int       `json:"level"`
	CurrentExp int       `json:"current_exp"`
	ClassID    uuid.UUID `gorm:"type:uuid" json:"class_id"`
	CreatedAt  time.Time `json:"created_at"`
}
