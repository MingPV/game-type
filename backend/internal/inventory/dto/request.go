package dto

type CreateInventoryRequest struct {
	MaxSlots int `json:"max_slots"`
}

// 	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"inventory_id"`
// 	MaxSlots int       `json:"max_slots"`
