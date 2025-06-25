package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Rarity struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"rarity_id"`
	Name     string    `json:"name"`
	DropRate float64   `json:"drop_rate"`
}

func (r *Rarity) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.New()
	return
}
