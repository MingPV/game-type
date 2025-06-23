package repository

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	"gorm.io/gorm"
)

type GormMonsterTypeRepository struct {
	db *gorm.DB
}

func NewGormMonsterTypeRepository(db *gorm.DB) MonsterTypeRepository {
	return &GormMonsterTypeRepository{db: db}
}

func (r *GormMonsterTypeRepository) Save(monsterType *entities.MonsterType) error {
	return r.db.Create(&monsterType).Error
}

func (r *GormMonsterTypeRepository) FindAll() ([]*entities.MonsterType, error) {
	var monsterTypeValues []entities.MonsterType
	if err := r.db.Find(&monsterTypeValues).Error; err != nil {
		return nil, err
	}

	monsterTypes := make([]*entities.MonsterType, len(monsterTypeValues))
	for i := range monsterTypeValues {
		monsterTypes[i] = &monsterTypeValues[i]
	}
	return monsterTypes, nil
}

func (r *GormMonsterTypeRepository) FindByID(id string) (*entities.MonsterType, error) {
	var monsterType entities.MonsterType
	if err := r.db.Where("id = ?", id).First(&monsterType).Error; err != nil {
		return &entities.MonsterType{}, err
	}
	return &monsterType, nil
}

func (r *GormMonsterTypeRepository) Patch(id string, monsterType *entities.MonsterType) error {
	if err := r.db.Model(&entities.MonsterType{}).Where("id = ?", id).Updates(monsterType).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormMonsterTypeRepository) Delete(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&entities.MonsterType{}).Error; err != nil {
		return err
	}
	return nil
}
