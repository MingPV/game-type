package repository

import "github.com/MingPV/clean-go-template/internal/entities"

type EquipmentSlotRepository interface {
	Save(equipmentSlot *entities.EquipmentSlot) error
	FindAll() ([]*entities.EquipmentSlot, error)
	FindByID(id int) (*entities.EquipmentSlot, error)
	// Patch(id int, equipmentSlot *entities.EquipmentSlot) error
	Delete(id int) error
}
