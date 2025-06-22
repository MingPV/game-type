package repository

import "github.com/MingPV/clean-go-template/internal/entities"

type InventoryRepository interface {
	Save(inventory *entities.Inventory) error
	FindAll() ([]*entities.Inventory, error)
	FindByID(id int) (*entities.Inventory, error)
	// Patch(id int, inventory *entities.Inventory) error
	Delete(id int) error
}
