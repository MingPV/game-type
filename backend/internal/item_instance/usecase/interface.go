package usecase

import "github.com/MingPV/clean-go-template/internal/entities"

type ItemInstanceUseCase interface {
	FindAllItemInstances() ([]*entities.ItemInstance, error)
	CreateItemInstance(itemInstance *entities.ItemInstance) (*entities.ItemInstance, error)
	PatchItemInstance(id string, itemInstance *entities.ItemInstance) error
	DeleteItemInstance(id string) error
	FindItemInstanceByID(id string) (*entities.ItemInstance, error)
}
