package repository

import "github.com/MingPV/clean-go-template/internal/entities"

type LevelProgressRepository interface {
	Save(level_progress *entities.LevelProgress) error
	FindAll() ([]*entities.LevelProgress, error)
	FindByLevel(level int) (*entities.LevelProgress, error)
	Patch(id int, level_progress *entities.LevelProgress) error
	Delete(id int) error
}
