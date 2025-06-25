package entities

import (
	"time"

	"github.com/google/uuid"
)

type Character struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"character_id"`
	UserID      uuid.UUID `gorm:"type:uuid" json:"user_id"`
	Name        string    `gorm:"uniqueIndex" json:"name"`
	Level       int       `json:"level"`
	CurrentExp  int       `json:"current_exp"`
	ClassID     uuid.UUID `gorm:"type:uuid" json:"class_id"`
	StatusID    uuid.UUID `gorm:"type:uuid" json:"status_id"`
	InventoryID uuid.UUID `gorm:"type:uuid" json:"inventory_id"`
	CreatedAt   time.Time `json:"created_at"`

	User           User            `gorm:"foreignKey:UserID;references:ID"`                       // this.UserID -> User.ID
	Class          Class           `gorm:"foreignKey:ClassID;references:ID" json:"class"`         // this.ClassID -> Class.ID
	Status         Status          `gorm:"foreignKey:ID;references:CharacterID" json:"status"`    // this.StatusID -> Status.CharacterID
	EquipmentSlots []EquipmentSlot `gorm:"foreignKey:CharacterID" json:"equipment_slots"`         // Slots.CharacterID -> this.ID
	Inventory      Inventory       `gorm:"foreignKey:InventoryID;references:ID" json:"inventory"` // this.InventoryID -> Inventory.ID
}

// No need to auto generate UID because it need to be set together with base status

// func (c *Character) BeforeCreate(tx *gorm.DB) (err error) {
// 	c.ID = uuid.New()
// 	return
// }
