package repository

import "github.com/MingPV/clean-go-template/internal/entities"

type RarityRepository interface {
	Save(rarity *entities.Rarity) error
	FindAll() ([]*entities.Rarity, error)
	FindByID(id int) (*entities.Rarity, error)
	// Patch(id int, rarity *entities.Rarity) error
	Delete(id int) error
}
