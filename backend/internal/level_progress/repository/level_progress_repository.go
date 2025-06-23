package repository

import "github.com/MingPV/clean-go-template/internal/entities"

type LevelProgressRepository interface {
	Save(level_progress *entities.LevelProgress) error
	FindAll() ([]*entities.LevelProgress, error)
	FindByID(id string) (*entities.LevelProgress, error)
	Patch(id string, level_progress *entities.LevelProgress) error
	Delete(id string) error
}
