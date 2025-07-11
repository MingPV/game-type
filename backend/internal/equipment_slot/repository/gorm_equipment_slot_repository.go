package repository

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	"gorm.io/gorm"
)

type GormEquipmentSlotRepository struct {
	db *gorm.DB
}

func NewGormEquipmentSlotRepository(db *gorm.DB) EquipmentSlotRepository {
	return &GormEquipmentSlotRepository{db: db}
}

func (r *GormEquipmentSlotRepository) Save(equipmentSlot *entities.EquipmentSlot) error {
	return r.db.Create(&equipmentSlot).Error
}

func (r *GormEquipmentSlotRepository) FindAll() ([]*entities.EquipmentSlot, error) {
	var equipmentSlotValues []entities.EquipmentSlot
	if err := r.db.
		Preload("ItemInstance").
		Preload("ItemInstance.Item").
		Preload("ItemInstance.Item.Rarity").
		Preload("ItemInstance.Item.ItemType").
		Preload("ItemInstance.Item.ItemStats").
		Find(&equipmentSlotValues).Error; err != nil {
		return nil, err
	}

	equipmentSlots := make([]*entities.EquipmentSlot, len(equipmentSlotValues))
	for i := range equipmentSlotValues {
		equipmentSlots[i] = &equipmentSlotValues[i]
	}
	return equipmentSlots, nil
}

func (r *GormEquipmentSlotRepository) FindByID(id string) (*entities.EquipmentSlot, error) {
	var equipmentSlot entities.EquipmentSlot
	if err := r.db.
		Preload("ItemInstance").
		Preload("ItemInstance.Item").
		Preload("ItemInstance.Item.Rarity").
		Preload("ItemInstance.Item.ItemType").
		Preload("ItemInstance.Item.ItemStats").
		Where("id = ?", id).
		First(&equipmentSlot).Error; err != nil {
		return &entities.EquipmentSlot{}, err
	}
	return &equipmentSlot, nil
}

func (r *GormEquipmentSlotRepository) Patch(id string, equipmentSlot *entities.EquipmentSlot) error {
	if err := r.db.Model(&entities.EquipmentSlot{}).Where("id = ?", id).Updates(equipmentSlot).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormEquipmentSlotRepository) Delete(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&entities.EquipmentSlot{}).Error; err != nil {
		return err
	}
	return nil
}
