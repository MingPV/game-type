package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Inventory struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"inventory_id"`
	MaxSlots int       `json:"max_slots"`

	ItemInstances []ItemInstance `gorm:"foreignKey:InventoryID;references:ID" json:"item_instances"` // ItemInstance.InventoryID -> this.ID
}

func (i *Inventory) BeforeCreate(tx *gorm.DB) (err error) {
	i.ID = uuid.New()
	return
}
