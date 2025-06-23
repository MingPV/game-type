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
	if err := r.db.Find(&monsterLootValues).Error; err != nil {
		return nil, err
	}

	monsterLoots := make([]*entities.MonsterLoot, len(monsterLootValues))
	for i := range monsterLootValues {
		monsterLoots[i] = &monsterLootValues[i]
	}
	return monsterLoots, nil
}

func (r *GormMonsterLootRepository) FindByID(id string) (*entities.MonsterLoot, error) {
	var monsterLoot entities.MonsterLoot
	if err := r.db.Where("id = ?", id).First(&monsterLoot).Error; err != nil {
		return &entities.MonsterLoot{}, err
	}
	return &monsterLoot, nil
}

func (r *GormMonsterLootRepository) Patch(id string, monsterLoot *entities.MonsterLoot) error {
	if err := r.db.Model(&entities.MonsterLoot{}).Where("id = ?", id).Updates(monsterLoot).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormMonsterLootRepository) Delete(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&entities.MonsterLoot{}).Error; err != nil {
		return err
	}
	return nil
}
