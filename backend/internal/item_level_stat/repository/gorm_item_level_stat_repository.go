package repository

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	"gorm.io/gorm"
)

type GormItemLevelStatRepository struct {
	db *gorm.DB
}

func NewGormItemLevelStatRepository(db *gorm.DB) ItemLevelStatRepository {
	return &GormItemLevelStatRepository{db: db}
}

func (r *GormItemLevelStatRepository) Save(itemLevelStat *entities.ItemLevelStat) error {
	return r.db.Create(&itemLevelStat).Error
}

func (r *GormItemLevelStatRepository) FindAll() ([]*entities.ItemLevelStat, error) {
	var itemLevelStatValues []entities.ItemLevelStat
	if err := r.db.Find(&itemLevelStatValues).Error; err != nil {
		return nil, err
	}

	itemLevelStats := make([]*entities.ItemLevelStat, len(itemLevelStatValues))
	for i := range itemLevelStatValues {
		itemLevelStats[i] = &itemLevelStatValues[i]
	}
	return itemLevelStats, nil
}

func (r *GormItemLevelStatRepository) FindByID(id string) (*entities.ItemLevelStat, error) {
	var itemLevelStat entities.ItemLevelStat
	if err := r.db.Where("id = ?", id).First(&itemLevelStat).Error; err != nil {
		return &entities.ItemLevelStat{}, err
	}
	return &itemLevelStat, nil
}

func (r *GormItemLevelStatRepository) Patch(id string, itemLevelStat *entities.ItemLevelStat) error {
	if err := r.db.Model(&entities.ItemLevelStat{}).Where("id = ?", id).Updates(itemLevelStat).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormItemLevelStatRepository) Delete(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&entities.ItemLevelStat{}).Error; err != nil {
		return err
	}
	return nil
}
