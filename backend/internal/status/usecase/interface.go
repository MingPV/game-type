package usecase

import "github.com/MingPV/clean-go-template/internal/entities"

type StatusUseCase interface {
	FindAllStatuses() ([]*entities.Status, error)
	CreateStatus(status *entities.Status) error
	// PatchStatus(id int, status *entities.Status) error
	DeleteStatus(id int) error
	FindStatusByID(id int) (*entities.Status, error)
}
