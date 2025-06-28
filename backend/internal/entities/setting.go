package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Setting struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"setting_id"`
	UserID      uuid.UUID `gorm:"type:uuid;uniqueIndex" json:"user_id"`
	MusicVolume float64   `json:"music_volume"`
	Language    string    `json:"language"`
}

func (s *Setting) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New()
	return
}

// Tested
