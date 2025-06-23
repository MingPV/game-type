package usecase

import "github.com/MingPV/clean-go-template/internal/entities"

type EquipmentSlotUseCase interface {
	FindAllEquipmentSlots() ([]*entities.EquipmentSlot, error)
	CreateEquipmentSlot(equipmentSlot *entities.EquipmentSlot) error
	PatchEquipmentSlot(id string, equipmentSlot *entities.EquipmentSlot) error
	DeleteEquipmentSlot(id string) error
	FindEquipmentSlotByID(id string) (*entities.EquipmentSlot, error)
}
