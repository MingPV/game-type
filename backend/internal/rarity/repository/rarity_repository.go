package repository

import "github.com/MingPV/clean-go-template/internal/entities"

type RarityRepository interface {
	Save(rarity *entities.Rarity) error
	FindAll() ([]*entities.Rarity, error)
	FindByID(id string) (*entities.Rarity, error)
	Patch(id string, rarity *entities.Rarity) error
	Delete(id string) error
}
