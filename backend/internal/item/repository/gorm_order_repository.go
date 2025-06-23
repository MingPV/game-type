package repository

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	"gorm.io/gorm"
)

type GormItemRepository struct {
	db *gorm.DB
}

func NewGormItemRepository(db *gorm.DB) ItemRepository {
	return &GormItemRepository{db: db}
}

func (r *GormItemRepository) Save(item *entities.Item) error {
	return r.db.Create(&item).Error
}

func (r *GormItemRepository) FindAll() ([]*entities.Item, error) {
	var itemValues []entities.Item
	if err := r.db.Find(&itemValues).Error; err != nil {
		return nil, err
	}

	items := make([]*entities.Item, len(itemValues))
	for i := range itemValues {
		items[i] = &itemValues[i]
	}
	return items, nil
}

func (r *GormItemRepository) FindByID(id string) (*entities.Item, error) {
	var item entities.Item
	if err := r.db.Where("id = ?", id).First(&item).Error; err != nil {
		return &entities.Item{}, err
	}
	return &item, nil
}

func (r *GormItemRepository) Patch(id string, item *entities.Item) error {
	if err := r.db.Model(&entities.Item{}).Where("id = ?", id).Updates(item).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormItemRepository) Delete(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&entities.Item{}).Error; err != nil {
		return err
	}
	return nil
}
