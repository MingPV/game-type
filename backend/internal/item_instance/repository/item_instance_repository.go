package repository

import "github.com/MingPV/clean-go-template/internal/entities"

type ItemInstanceRepository interface {
	Save(itemInstance *entities.ItemInstance) error
	FindAll() ([]*entities.ItemInstance, error)
	FindByID(id string) (*entities.ItemInstance, error)
	Patch(id string, itemInstance *entities.ItemInstance) error
	Delete(id string) error
}
