package repository

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	"gorm.io/gorm"
)

type GormItemInstanceRepository struct {
	db *gorm.DB
}

func NewGormItemInstanceRepository(db *gorm.DB) ItemInstanceRepository {
	return &GormItemInstanceRepository{db: db}
}

func (r *GormItemInstanceRepository) Save(itemInstance *entities.ItemInstance) error {
	return r.db.Create(&itemInstance).Error
}

func (r *GormItemInstanceRepository) FindAll() ([]*entities.ItemInstance, error) {
	var itemInstanceValues []entities.ItemInstance
	if err := r.db.Preload("Item").Preload("Item.ItemType").Preload("Item.Rarity").Preload("Item.ItemStats").Find(&itemInstanceValues).Error; err != nil {
		return nil, err
	}

	itemInstances := make([]*entities.ItemInstance, len(itemInstanceValues))
	for i := range itemInstanceValues {
		itemInstances[i] = &itemInstanceValues[i]
	}
	return itemInstances, nil
}

func (r *GormItemInstanceRepository) FindByID(id string) (*entities.ItemInstance, error) {
	var itemInstance entities.ItemInstance
	if err := r.db.Preload("Item").Preload("Item.ItemType").Preload("Item.Rarity").Preload("Item.ItemStats").Where("id = ?", id).First(&itemInstance).Error; err != nil {
		return &entities.ItemInstance{}, err
	}
	return &itemInstance, nil
}

func (r *GormItemInstanceRepository) Patch(id string, itemInstance *entities.ItemInstance) error {
	if err := r.db.Model(&entities.ItemInstance{}).Where("id = ?", id).Updates(itemInstance).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormItemInstanceRepository) Delete(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&entities.ItemInstance{}).Error; err != nil {
		return err
	}
	return nil
}
