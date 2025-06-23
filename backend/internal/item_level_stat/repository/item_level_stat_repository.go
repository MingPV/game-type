package repository

import "github.com/MingPV/clean-go-template/internal/entities"

type ItemLevelStatRepository interface {
	Save(itemLevelStat *entities.ItemLevelStat) error
	FindAll() ([]*entities.ItemLevelStat, error)
	FindByID(id string) (*entities.ItemLevelStat, error)
	// Patch(id int, itemLevelStat *entities.ItemLevelStat) error
	Delete(id int) error
}
