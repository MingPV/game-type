package repository

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	"gorm.io/gorm"
)

type GormMonsterRepository struct {
	db *gorm.DB
}

func NewGormMonsterRepository(db *gorm.DB) MonsterRepository {
	return &GormMonsterRepository{db: db}
}

func (r *GormMonsterRepository) Save(monster *entities.Monster) error {
	return r.db.Create(&monster).Error
}

func (r *GormMonsterRepository) FindAll() ([]*entities.Monster, error) {
	var monsterValues []entities.Monster
	if err := r.db.
		Preload("MonsterType").
		Preload("MonsterLoots").
		Preload("MonsterLoots.Item").
		Preload("MonsterLoots.Item.ItemType").
		Preload("MonsterLoots.Item.Rarity").
		Preload("MonsterLoots.Item.ItemStats").
		Find(&monsterValues).Error; err != nil {
		return nil, err
	}

	monsters := make([]*entities.Monster, len(monsterValues))
	for i := range monsterValues {
		monsters[i] = &monsterValues[i]
	}
	return monsters, nil
}

func (r *GormMonsterRepository) FindByID(id string) (*entities.Monster, error) {
	var monster entities.Monster
	if err := r.db.
		Preload("MonsterType").
		Preload("MonsterLoots").
		Preload("MonsterLoots.Item").
		Preload("MonsterLoots.Item.ItemType").
		Preload("MonsterLoots.Item.Rarity").
		Preload("MonsterLoots.Item.ItemStats").
		Where("id = ?", id).
		First(&monster).Error; err != nil {
		return &entities.Monster{}, err
	}
	return &monster, nil
}

func (r *GormMonsterRepository) Patch(id string, monster *entities.Monster) error {
	if err := r.db.Model(&entities.Monster{}).Where("id = ?", id).Updates(monster).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormMonsterRepository) Delete(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&entities.Monster{}).Error; err != nil {
		return err
	}
	return nil
}
