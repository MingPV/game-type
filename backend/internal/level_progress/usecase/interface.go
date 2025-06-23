package usecase

import "github.com/MingPV/clean-go-template/internal/entities"

type LevelProgressUseCase interface {
	FindAllLevelProgresses() ([]*entities.LevelProgress, error)
	CreateLevelProgress(level_progress *entities.LevelProgress) error
	PatchLevelProgress(id string, level_progress *entities.LevelProgress) error
	DeleteLevelProgress(id string) error
	FindLevelProgressByID(id string) (*entities.LevelProgress, error)
}
