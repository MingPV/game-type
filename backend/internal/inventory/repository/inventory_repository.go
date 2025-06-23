package repository

import "github.com/MingPV/clean-go-template/internal/entities"

type InventoryRepository interface {
	Save(inventory *entities.Inventory) error
	FindAll() ([]*entities.Inventory, error)
	FindByID(id string) (*entities.Inventory, error)
	Patch(id string, inventory *entities.Inventory) error
	Delete(id string) error
}
