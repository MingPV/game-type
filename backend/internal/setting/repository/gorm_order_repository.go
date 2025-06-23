package repository

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	"gorm.io/gorm"
)

type GormSettingRepository struct {
	db *gorm.DB
}

func NewGormSettingRepository(db *gorm.DB) SettingRepository {
	return &GormSettingRepository{db: db}
}

func (r *GormSettingRepository) Save(setting *entities.Setting) error {
	return r.db.Create(&setting).Error
}

func (r *GormSettingRepository) FindAll() ([]*entities.Setting, error) {
	var settingValues []entities.Setting
	if err := r.db.Find(&settingValues).Error; err != nil {
		return nil, err
	}

	settings := make([]*entities.Setting, len(settingValues))
	for i := range settingValues {
		settings[i] = &settingValues[i]
	}
	return settings, nil
}

func (r *GormSettingRepository) FindByID(id string) (*entities.Setting, error) {
	var setting entities.Setting
	if err := r.db.Where("id = ?", id).First(&setting).Error; err != nil {
		return &entities.Setting{}, err
	}
	return &setting, nil
}

func (r *GormSettingRepository) Patch(id string, setting *entities.Setting) error {
	if err := r.db.Model(&entities.Setting{}).Where("id = ?", id).Updates(setting).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormSettingRepository) Delete(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&entities.Setting{}).Error; err != nil {
		return err
	}
	return nil
}
