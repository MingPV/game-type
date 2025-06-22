package usecase

import "github.com/MingPV/clean-go-template/internal/entities"

type ClassUseCase interface {
	FindAllClasses() ([]*entities.Class, error)
	CreateClass(class *entities.Class) error
	// PatchClass(id int, class *entities.Class) error
	DeleteClass(id int) error
	FindClassByID(id int) (*entities.Class, error)
}
