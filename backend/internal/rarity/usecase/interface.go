package usecase

import "github.com/MingPV/clean-go-template/internal/entities"

type RarityUseCase interface {
	FindAllRarities() ([]*entities.Rarity, error)
	CreateRarity(rarity *entities.Rarity) error
	// PatchRarity(id int, rarity *entities.Rarity) error
	DeleteRarity(id int) error
	FindRarityByID(id string) (*entities.Rarity, error)
}
