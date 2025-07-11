package repository

import "github.com/MingPV/clean-go-template/internal/entities"

type ItemTypeRepository interface {
	Save(itemType *entities.ItemType) error
	FindAll() ([]*entities.ItemType, error)
	FindByID(id string) (*entities.ItemType, error)
	Patch(id string, itemType *entities.ItemType) error
	Delete(id string) error
}
