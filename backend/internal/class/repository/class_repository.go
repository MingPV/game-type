package repository

import "github.com/MingPV/clean-go-template/internal/entities"

type ClassRepository interface {
	Save(class *entities.Class) error
	FindAll() ([]*entities.Class, error)
	FindByID(id int) (*entities.Class, error)
	// Patch(id int, class *entities.Class) error
	Delete(id int) error
}
