package usecase

import "github.com/MingPV/clean-go-template/internal/entities"

type ItemInstanceUseCase interface {
	FindAllItemInstances() ([]*entities.ItemInstance, error)
	CreateItemInstance(itemInstance *entities.ItemInstance) error
	// PatchItemInstance(id int, itemInstance *entities.ItemInstance) error
	DeleteItemInstance(id int) error
	FindItemInstanceByID(id string) (*entities.ItemInstance, error)
}
