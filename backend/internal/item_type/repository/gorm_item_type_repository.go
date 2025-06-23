package repository

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	"gorm.io/gorm"
)

type GormItemTypeRepository struct {
	db *gorm.DB
}

func NewGormItemTypeRepository(db *gorm.DB) ItemTypeRepository {
	return &GormItemTypeRepository{db: db}
}

func (r *GormItemTypeRepository) Save(itemType *entities.ItemType) error {
	return r.db.Create(&itemType).Error
}

func (r *GormItemTypeRepository) FindAll() ([]*entities.ItemType, error) {
	var itemTypeValues []entities.ItemType
	if err := r.db.Find(&itemTypeValues).Error; err != nil {
		return nil, err
	}

	itemTypes := make([]*entities.ItemType, len(itemTypeValues))
	for i := range itemTypeValues {
		itemTypes[i] = &itemTypeValues[i]
	}
	return itemTypes, nil
}

func (r *GormItemTypeRepository) FindByID(id string) (*entities.ItemType, error) {
	var itemType entities.ItemType
	if err := r.db.Where("id = ?", id).First(&itemType).Error; err != nil {
		return &entities.ItemType{}, err
	}
	return &itemType, nil
}

func (r *GormItemTypeRepository) Patch(id string, itemType *entities.ItemType) error {
	if err := r.db.Model(&entities.ItemType{}).Where("id = ?", id).Updates(itemType).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormItemTypeRepository) Delete(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&entities.ItemType{}).Error; err != nil {
		return err
	}
	return nil
}
