package repository

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	"gorm.io/gorm"
)

type GormStatusRepository struct {
	db *gorm.DB
}

func NewGormStatusRepository(db *gorm.DB) StatusRepository {
	return &GormStatusRepository{db: db}
}

func (r *GormStatusRepository) Save(status *entities.Status) error {
	return r.db.Create(&status).Error
}

func (r *GormStatusRepository) FindAll() ([]*entities.Status, error) {
	var statusValues []entities.Status
	if err := r.db.Find(&statusValues).Error; err != nil {
		return nil, err
	}

	statuses := make([]*entities.Status, len(statusValues))
	for i := range statusValues {
		statuses[i] = &statusValues[i]
	}
	return statuses, nil
}

func (r *GormStatusRepository) FindByCharacterID(character_id string) (*entities.Status, error) {
	var status entities.Status
	if err := r.db.Where("character_id = ?", character_id).First(&status).Error; err != nil {
		return &entities.Status{}, err
	}
	return &status, nil
}

func (r *GormStatusRepository) Patch(character_id string, status *entities.Status) error {
	if err := r.db.Model(&entities.Status{}).Where("character_id = ?", character_id).Updates(status).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormStatusRepository) Delete(character_id string) error {
	if err := r.db.Where("character_id = ?", character_id).Delete(&entities.Status{}).Error; err != nil {
		return err
	}
	return nil
}
