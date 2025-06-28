package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ItemInstance struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"item_instance_id"`
	InventoryID  uuid.UUID `gorm:"type:uuid" json:"inventory_id"`
	ItemID       uuid.UUID `gorm:"type:uuid" json:"item_id"`
	UpgradeLevel int       `json:"upgrade_level"`
	CreatedAt    time.Time `json:"created_at"`

	Item Item `gorm:"foreignKey:ItemID;references:ID" json:"item"` // this.ItemID -> Item.ID
}

func (i *ItemInstance) BeforeCreate(tx *gorm.DB) (err error) {
	i.ID = uuid.New()
	return
}

// Tested
