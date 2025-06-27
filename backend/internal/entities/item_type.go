package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ItemType struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"item_type_id"`
	Name string    `json:"name"`
}

func (i *ItemType) BeforeCreate(tx *gorm.DB) (err error) {
	i.ID = uuid.New()
	return
}

// Tested
