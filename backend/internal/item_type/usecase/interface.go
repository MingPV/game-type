package usecase

import "github.com/MingPV/clean-go-template/internal/entities"

type ItemTypeUseCase interface {
	FindAllItemTypes() ([]*entities.ItemType, error)
	CreateItemType(itemType *entities.ItemType) error
	PatchItemType(id string, itemType *entities.ItemType) error
	DeleteItemType(id string) error
	FindItemTypeByID(id string) (*entities.ItemType, error)
}
