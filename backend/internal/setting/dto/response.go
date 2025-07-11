package dto

import (
	"github.com/google/uuid"
)

type SettingResponse struct {
	ID          uuid.UUID `gorm:"type:uuid" json:"setting_id"`
	UserID      uuid.UUID `gorm:"type:uuid" json:"user_id"`
	MusicVolume float64   `json:"music_volume"`
	Language    string    `json:"language"`
}
