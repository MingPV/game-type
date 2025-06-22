package dto

import (
	"time"

	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/google/uuid"
)

type CharacterResponse struct {
	ID         uuid.UUID `gorm:"type:uuid" json:"character_id"`
	UserID     uuid.UUID `gorm:"type:uuid" json:"user_id"`
	Name       string    `json:"name"`
	Level      int       `json:"level"`
	CurrentExp int       `json:"current_exp"`
	ClassID    uuid.UUID `gorm:"type:uuid" json:"class_id"`
	CreatedAt  time.Time `json:"created_at"`

	Class  entities.Class           `json:"class"`
	Status entities.Status          `json:"status"`
	Slots  []entities.EquipmentSlot `json:"equipment_slots"`
	Items  []entities.ItemInstance  `json:"items"`
}

// type Character struct {
// 	ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"character_id"`
// 	UserID     uuid.UUID `gorm:"type:uuid" json:"user_id"`
// 	Name       string    `json:"name"`
// 	Level      int       `json:"level"`
// 	CurrentExp int       `json:"current_exp"`
// 	ClassID    uuid.UUID `gorm:"type:uuid" json:"class_id"`
// 	StatusID   uuid.UUID `gorm:"type:uuid" json:"status_id"`
// 	CreatedAt  time.Time `json:"created_at"`

// 	User   User            `gorm:"foreignKey:UserID;references:ID"`
// 	Class  Class           `gorm:"foreignKey:ClassID;references:ID" json:"class"`
// 	Status Status          `gorm:"foreignKey:StatusID;references:CharacterID" json:"status"`
// 	Slots  []EquipmentSlot `gorm:"foreignKey:CharacterID" json:"equipment_slots"`
// 	Items  []ItemInstance  `gorm:"foreignKey:OwnerCharacterID" json:"items"`
// }
