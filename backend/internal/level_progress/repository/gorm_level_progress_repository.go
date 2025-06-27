package repository

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	"gorm.io/gorm"
)

type GormLevelProgressRepository struct {
	db *gorm.DB
}

func NewGormLevelProgressRepository(db *gorm.DB) LevelProgressRepository {
	return &GormLevelProgressRepository{db: db}
}

func (r *GormLevelProgressRepository) Save(level_progress *entities.LevelProgress) error {
	return r.db.Create(&level_progress).Error
}

func (r *GormLevelProgressRepository) FindAll() ([]*entities.LevelProgress, error) {
	var level_progressValues []entities.LevelProgress
	if err := r.db.Find(&level_progressValues).Error; err != nil {
		return nil, err
	}

	level_progresses := make([]*entities.LevelProgress, len(level_progressValues))
	for i := range level_progressValues {
		level_progresses[i] = &level_progressValues[i]
	}
	return level_progresses, nil
}

func (r *GormLevelProgressRepository) FindByLevel(level int) (*entities.LevelProgress, error) {
	var level_progress entities.LevelProgress
	if err := r.db.Where("level = ?", level).First(&level_progress).Error; err != nil {
		return &entities.LevelProgress{}, err
	}
	return &level_progress, nil
}

func (r *GormLevelProgressRepository) Patch(level int, level_progress *entities.LevelProgress) error {
	if err := r.db.Model(&entities.LevelProgress{}).Where("level = ?", level).Updates(level_progress).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormLevelProgressRepository) Delete(level int) error {
	if err := r.db.Where("level = ?", level).Delete(&entities.LevelProgress{}).Error; err != nil {
		return err
	}
	return nil
}
