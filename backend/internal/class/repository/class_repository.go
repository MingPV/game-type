package repository

import "github.com/MingPV/clean-go-template/internal/entities"

type ClassRepository interface {
	Save(class *entities.Class) error
	FindAll() ([]*entities.Class, error)
	FindByID(id string) (*entities.Class, error)
	Patch(id string, class *entities.Class) error
	Delete(id string) error
}
