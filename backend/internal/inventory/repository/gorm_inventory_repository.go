package repository

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	"gorm.io/gorm"
)

type GormInventoryRepository struct {
	db *gorm.DB
}

func NewGormInventoryRepository(db *gorm.DB) InventoryRepository {
	return &GormInventoryRepository{db: db}
}

func (r *GormInventoryRepository) Save(inventory *entities.Inventory) error {
	return r.db.Create(&inventory).Error
}

func (r *GormInventoryRepository) FindAll() ([]*entities.Inventory, error) {
	var inventoryValues []entities.Inventory
	if err := r.db.
		Preload("ItemInstances").
		Preload("ItemInstances.Item").
		Preload("ItemInstances.Item.Rarity").
		Preload("ItemInstances.Item.ItemType").
		Preload("ItemInstances.Item.ItemStats").
		Find(&inventoryValues).Error; err != nil {
		return nil, err
	}

	inventories := make([]*entities.Inventory, len(inventoryValues))
	for i := range inventoryValues {
		inventories[i] = &inventoryValues[i]
	}
	return inventories, nil
}

func (r *GormInventoryRepository) FindByID(id string) (*entities.Inventory, error) {
	var inventory entities.Inventory
	if err := r.db.Where("id = ?", id).
		Preload("ItemInstances").
		Preload("ItemInstances.Item").
		Preload("ItemInstances.Item.Rarity").
		Preload("ItemInstances.Item.ItemType").
		Preload("ItemInstances.Item.ItemStats").
		First(&inventory).Error; err != nil {
		return &entities.Inventory{}, err
	}
	return &inventory, nil
}

func (r *GormInventoryRepository) Patch(id string, inventory *entities.Inventory) error {
	if err := r.db.Model(&entities.Inventory{}).Where("id = ?", id).Updates(inventory).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormInventoryRepository) Delete(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&entities.Inventory{}).Error; err != nil {
		return err
	}
	return nil
}
