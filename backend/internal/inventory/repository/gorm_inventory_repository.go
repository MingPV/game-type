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
	if err := r.db.Find(&inventoryValues).Error; err != nil {
		return nil, err
	}

	inventorys := make([]*entities.Inventory, len(inventoryValues))
	for i := range inventoryValues {
		inventorys[i] = &inventoryValues[i]
	}
	return inventorys, nil
}

func (r *GormInventoryRepository) FindByID(id int) (*entities.Inventory, error) {
	var inventory entities.Inventory
	if err := r.db.First(&inventory, id).Error; err != nil {
		return &entities.Inventory{}, err
	}
	return &inventory, nil
}

// func (r *GormInventoryRepository) Patch(id int, inventory *entities.Inventory) error {
// 	if err := r.db.Model(&entities.Inventory{}).Where("id = ?", id).Updates(inventory).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

func (r *GormInventoryRepository) Delete(id int) error {
	if err := r.db.Delete(&entities.Inventory{}, id).Error; err != nil {
		return err
	}
	return nil
}
