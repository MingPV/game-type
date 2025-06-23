package usecase

import "github.com/MingPV/clean-go-template/internal/entities"

type ItemLevelStatUseCase interface {
	FindAllItemLevelStats() ([]*entities.ItemLevelStat, error)
	CreateItemLevelStat(itemLevelStat *entities.ItemLevelStat) error
	PatchItemLevelStat(id string, itemLevelStat *entities.ItemLevelStat) error
	DeleteItemLevelStat(id string) error
	FindItemLevelStatByID(id string) (*entities.ItemLevelStat, error)
}
