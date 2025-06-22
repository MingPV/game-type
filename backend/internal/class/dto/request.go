package dto

import (
	"github.com/google/uuid"
)

type CreateClassRequest struct {
	ID          uuid.UUID `gorm:"type:uuid" json:"class_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}
