package usecase

import "github.com/MingPV/clean-go-template/internal/entities"

type LevelProgressUseCase interface {
	FindAllLevelProgresses() ([]*entities.LevelProgress, error)
	CreateLevelProgress(level_progress *entities.LevelProgress) error
	PatchLevelProgress(id int, level_progress *entities.LevelProgress) error
	DeleteLevelProgress(id int) error
	FindLevelProgressByLevel(level int) (*entities.LevelProgress, error)
}
