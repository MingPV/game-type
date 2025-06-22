package repository

import "github.com/MingPV/clean-go-template/internal/entities"

type StatusRepository interface {
	Save(status *entities.Status) error
	FindAll() ([]*entities.Status, error)
	FindByID(id int) (*entities.Status, error)
	// Patch(id int, status *entities.Status) error
	Delete(id int) error
}
