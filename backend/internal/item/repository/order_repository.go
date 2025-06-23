package repository

import "github.com/MingPV/clean-go-template/internal/entities"

type ItemRepository interface {
	Save(item *entities.Item) error
	FindAll() ([]*entities.Item, error)
	FindByID(id string) (*entities.Item, error)
	// Patch(id int, item *entities.Item) error
	Delete(id int) error
}
