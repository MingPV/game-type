package usecase

import "github.com/MingPV/clean-go-template/internal/entities"

type EquipmentSlotUseCase interface {
	FindAllEquipmentSlots() ([]*entities.EquipmentSlot, error)
	CreateEquipmentSlot(equipmentSlot *entities.EquipmentSlot) error
	// PatchEquipmentSlot(id int, equipmentSlot *entities.EquipmentSlot) error
	DeleteEquipmentSlot(id int) error
	FindEquipmentSlotByID(id int) (*entities.EquipmentSlot, error)
}
