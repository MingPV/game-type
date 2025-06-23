package usecase

import "github.com/MingPV/clean-go-template/internal/entities"

type StatusUseCase interface {
	FindAllStatuses() ([]*entities.Status, error)
	CreateStatus(status *entities.Status) error
	PatchStatus(id string, status *entities.Status) error
	DeleteStatus(id string) error
	FindStatusByID(id string) (*entities.Status, error)
}
