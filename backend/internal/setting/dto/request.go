package dto

import "github.com/google/uuid"

type CreateSettingRequest struct {
	UserID      uuid.UUID `gorm:"type:uuid" json:"user_id"`
	MusicVolume float64   `json:"music_volume"`
	Language    string    `json:"language"`
}
