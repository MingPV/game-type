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

func (r *GormStatusRepository) FindByID(id int) (*entities.Status, error) {
	var status entities.Status
	if err := r.db.First(&status, id).Error; err != nil {
		return &entities.Status{}, err
	}
	return &status, nil
}

func (r *GormStatusRepository) Patch(id int, status *entities.Status) error {
	if err := r.db.Model(&entities.Status{}).Where("id = ?", id).Updates(status).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormStatusRepository) Delete(id int) error {
	if err := r.db.Delete(&entities.Status{}, id).Error; err != nil {
		return err
	}
	return nil
}
