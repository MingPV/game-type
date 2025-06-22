package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Inventory struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey" json:"inventory_id"`
	ItemInstanceID uuid.UUID `gorm:"type:uuid" json:"item_instance_id"`
	Quantity       int       `json:"quantity"`

	ItemInstance ItemInstance `gorm:"foreignKey:ItemInstanceID;references:ID" json:"item_instance"` // this.ItemInstanceID -> ItemInstance.ID
}

func (i *Inventory) BeforeCreate(tx *gorm.DB) (err error) {
	i.ID = uuid.New()
	return
}
