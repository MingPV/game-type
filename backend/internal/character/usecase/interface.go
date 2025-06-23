package usecase

import "github.com/MingPV/clean-go-template/internal/entities"

type CharacterUseCase interface {
	FindAllCharacters() ([]*entities.Character, error)
	CreateCharacter(character *entities.Character) (*entities.Character, error)
	// PatchCharacter(id int, character *entities.Character) error
	DeleteCharacter(id string) error
	FindCharacterByID(id string) (*entities.Character, error)
}
