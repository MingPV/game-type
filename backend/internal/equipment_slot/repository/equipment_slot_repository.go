package repository

import "github.com/MingPV/clean-go-template/internal/entities"

type EquipmentSlotRepository interface {
	Save(equipmentSlot *entities.EquipmentSlot) error
	FindAll() ([]*entities.EquipmentSlot, error)
	FindByID(id string) (*entities.EquipmentSlot, error)
	Patch(id string, equipmentSlot *entities.EquipmentSlot) error
	Delete(id string) error
}
