package usecase

import "github.com/MingPV/clean-go-template/internal/entities"

type ClassUseCase interface {
	FindAllClasses() ([]*entities.Class, error)
	CreateClass(class *entities.Class) error
	PatchClass(id string, class *entities.Class) error
	DeleteClass(id string) error
	FindClassByID(id string) (*entities.Class, error)
}
