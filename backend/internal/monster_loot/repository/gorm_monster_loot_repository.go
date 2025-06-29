package repository

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	"gorm.io/gorm"
)

type GormMonsterLootRepository struct {
	db *gorm.DB
}

func NewGormMonsterLootRepository(db *gorm.DB) MonsterLootRepository {
	return &GormMonsterLootRepository{db: db}
}

func (r *GormMonsterLootRepository) Save(monsterLoot *entities.MonsterLoot) error {
	return r.db.Create(&monsterLoot).Error
}

func (r *GormMonsterLootRepository) FindAll() ([]*entities.MonsterLoot, error) {
	var monsterLootValues []entities.MonsterLoot
	if err := r.db.
		Preload("Item").
		Preload("Item.ItemType").
		Preload("Item.Rarity").
		Preload("Item.ItemStats").
		Find(&monsterLootValues).Error; err != nil {
		return nil, err
	}

	monsterLoots := make([]*entities.MonsterLoot, len(monsterLootValues))
	for i := range monsterLootValues {
		monsterLoots[i] = &monsterLootValues[i]
	}
	return monsterLoots, nil
}

func (r *GormMonsterLootRepository) FindByMonsterID(monster_id string) ([]*entities.MonsterLoot, error) {
	var monsterLootValues []entities.MonsterLoot
	if err := r.db.
		Preload("Item").
		Preload("Item.ItemType").
		Preload("Item.Rarity").
		Preload("Item.ItemStats").
		Where("monster_id = ?", monster_id).
		Find(&monsterLootValues).Error; err != nil {
		return nil, err
	}

	monsterLoots := make([]*entities.MonsterLoot, len(monsterLootValues))
	for i := range monsterLootValues {
		monsterLoots[i] = &monsterLootValues[i]
	}
	return monsterLoots, nil
}

func (r *GormMonsterLootRepository) FindByItemID(item_id string) ([]*entities.MonsterLoot, error) {
	var monsterLootValues []entities.MonsterLoot
	if err := r.db.
		Preload("Item").
		Preload("Item.ItemType").
		Preload("Item.Rarity").
		Preload("Item.ItemStats").
		Where("item_id = ?", item_id).
		Find(&monsterLootValues).Error; err != nil {
		return nil, err
	}

	monsterLoots := make([]*entities.MonsterLoot, len(monsterLootValues))
	for i := range monsterLootValues {
		monsterLoots[i] = &monsterLootValues[i]
	}
	return monsterLoots, nil
}

func (r *GormMonsterLootRepository) FindByMonsterIDAndItemID(monsterID string, itemID string) (*entities.MonsterLoot, error) {
	var monsterLoot entities.MonsterLoot
	if err := r.db.
		Preload("Item").
		Preload("Item.ItemType").
		Preload("Item.Rarity").
		Preload("Item.ItemStats").
		Where("monster_id = ? AND item_id = ?", monsterID, itemID).
		First(&monsterLoot).Error; err != nil {
		return nil, err
	}
	return &monsterLoot, nil
}

func (r *GormMonsterLootRepository) Patch(monster_id string, item_id string, monsterLoot *entities.MonsterLoot) error {
	if err := r.db.Model(&entities.MonsterLoot{}).
		Where("monster_id = ? AND item_id = ?", monster_id, item_id).
		Updates(monsterLoot).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormMonsterLootRepository) Delete(monster_id string, item_id string) error {
	if err := r.db.Where("monster_id = ? AND item_id = ?", monster_id, item_id).
		Delete(&entities.MonsterLoot{}).
		Error; err != nil {
		return err
	}
	return nil
}
