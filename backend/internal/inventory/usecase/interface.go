package usecase

import "github.com/MingPV/clean-go-template/internal/entities"

type InventoryUseCase interface {
	FindAllInventories() ([]*entities.Inventory, error)
	CreateInventory(inventory *entities.Inventory) error
	PatchInventory(id string, inventory *entities.Inventory) error
	DeleteInventory(id string) error
	FindInventoryByID(id string) (*entities.Inventory, error)
}
