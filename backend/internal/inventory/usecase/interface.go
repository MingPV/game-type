package usecase

import "github.com/MingPV/clean-go-template/internal/entities"

type InventoryUseCase interface {
	FindAllInventories() ([]*entities.Inventory, error)
	CreateInventory(inventory *entities.Inventory) error
	// PatchInventory(id int, inventory *entities.Inventory) error
	DeleteInventory(id int) error
	FindInventoryByID(id int) (*entities.Inventory, error)
}
