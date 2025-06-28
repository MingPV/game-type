package repository

import "github.com/MingPV/clean-go-template/internal/entities"

type StatusRepository interface {
	Save(status *entities.Status) error
	FindAll() ([]*entities.Status, error)
	FindByCharacterID(character_id string) (*entities.Status, error)
	Patch(character_id string, status *entities.Status) error
	Delete(character_id string) error
}
