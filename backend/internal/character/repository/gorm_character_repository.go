package repository

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	"gorm.io/gorm"
)

type GormCharacterRepository struct {
	db *gorm.DB
}

func NewGormCharacterRepository(db *gorm.DB) CharacterRepository {
	return &GormCharacterRepository{db: db}
}

func (r *GormCharacterRepository) Save(character *entities.Character) error {
	return r.db.Create(&character).Error
}

func (r *GormCharacterRepository) FindAll() ([]*entities.Character, error) {
	var characterValues []entities.Character
	if err := r.db.
		Preload("Class").
		Preload("Status").
		Preload("EquipmentSlots").
		// Preload("Inventory").
		// Preload("Inventory.ItemInstances").
		// Preload("Inventory.ItemInstances.Item.Rarity").
		// Preload("Inventory.ItemInstances.Item.ItemType").
		// Preload("Inventory.ItemInstances.Item.ItemStats").
		Preload("EquipmentSlots.ItemInstance").
		Preload("EquipmentSlots.ItemInstance.Item").
		Preload("EquipmentSlots.ItemInstance.Item.Rarity").
		Preload("EquipmentSlots.ItemInstance.Item.ItemType").
		Preload("EquipmentSlots.ItemInstance.Item.ItemStats").
		Find(&characterValues).Error; err != nil {
		return nil, err
	}

	characters := make([]*entities.Character, len(characterValues))
	for i := range characterValues {
		characters[i] = &characterValues[i]
	}
	return characters, nil
}

func (r *GormCharacterRepository) FindByUserID(user_id string) ([]*entities.Character, error) {
	var characterValues []entities.Character
	if err := r.db.
		Preload("Class").
		Preload("Status").
		Preload("EquipmentSlots").
		// Preload("Inventory").
		// Preload("Inventory.ItemInstances").
		// Preload("Inventory.ItemInstances.Item.Rarity").
		// Preload("Inventory.ItemInstances.Item.ItemType").
		// Preload("Inventory.ItemInstances.Item.ItemStats").
		Preload("EquipmentSlots.ItemInstance").
		Preload("EquipmentSlots.ItemInstance.Item").
		Preload("EquipmentSlots.ItemInstance.Item.Rarity").
		Preload("EquipmentSlots.ItemInstance.Item.ItemType").
		Preload("EquipmentSlots.ItemInstance.Item.ItemStats").
		Where("user_id = ?", user_id).
		Find(&characterValues).Error; err != nil {
		return nil, err
	}

	characters := make([]*entities.Character, len(characterValues))
	for i := range characterValues {
		characters[i] = &characterValues[i]
	}
	return characters, nil
}

func (r *GormCharacterRepository) FindByID(id string) (*entities.Character, error) {
	var character entities.Character
	if err := r.db.
		Preload("Class").
		Preload("Status").
		Preload("EquipmentSlots").
		// Preload("Inventory").
		// Preload("Inventory.ItemInstances").
		// Preload("Inventory.ItemInstances.Item.Rarity").
		// Preload("Inventory.ItemInstances.Item.ItemType").
		// Preload("Inventory.ItemInstances.Item.ItemStats").
		Preload("EquipmentSlots.ItemInstance").
		Preload("EquipmentSlots.ItemInstance.Item").
		Preload("EquipmentSlots.ItemInstance.Item.Rarity").
		Preload("EquipmentSlots.ItemInstance.Item.ItemType").
		Preload("EquipmentSlots.ItemInstance.Item.ItemStats").
		Where("id = ?", id).
		First(&character).Error; err != nil {
		return &entities.Character{}, err
	}
	return &character, nil
}

func (r *GormCharacterRepository) Patch(id string, character *entities.Character) error {
	if err := r.db.Model(&entities.Character{}).Where("id = ?", id).Updates(character).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormCharacterRepository) Delete(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&entities.Character{}).Error; err != nil {
		return err
	}
	return nil
}
