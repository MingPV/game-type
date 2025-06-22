package dto

import "github.com/google/uuid"

type ItemTypeResponse struct {
	ID   uuid.UUID `gorm:"type:uuid" json:"item_type_id"`
	Name string    `json:"name"`
}

// type ItemType struct {
// 	ID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"item_type_id"`
// 	Name string    `json:"name"`
// }
