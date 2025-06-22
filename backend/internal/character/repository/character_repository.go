package repository

import "github.com/MingPV/clean-go-template/internal/entities"

type CharacterRepository interface {
	Save(character *entities.Character) error
	FindAll() ([]*entities.Character, error)
	FindByID(id int) (*entities.Character, error)
	// Patch(id int, character *entities.Character) error
	Delete(id int) error
}
