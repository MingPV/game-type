package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Class struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"class_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func (c *Class) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}

// Tested
