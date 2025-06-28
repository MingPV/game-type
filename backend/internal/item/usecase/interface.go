package usecase

import "github.com/MingPV/clean-go-template/internal/entities"

type ItemUseCase interface {
	FindAllItems() ([]*entities.Item, error)
	CreateItem(item *entities.Item, item_level_stat *entities.ItemLevelStat) (*entities.Item, error)
	PatchItem(id string, item *entities.Item) error
	DeleteItem(id string) error
	FindItemByID(id string) (*entities.Item, error)
}
