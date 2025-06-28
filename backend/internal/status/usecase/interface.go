package usecase

import "github.com/MingPV/clean-go-template/internal/entities"

type StatusUseCase interface {
	FindAllStatuses() ([]*entities.Status, error)
	CreateStatus(status *entities.Status) error
	PatchStatus(character_id string, status *entities.Status) error
	DeleteStatus(character_id string) error
	FindStatusByCharacterID(character_id string) (*entities.Status, error)
}
