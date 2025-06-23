package repository

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	"gorm.io/gorm"
)

type GormRarityRepository struct {
	db *gorm.DB
}

func NewGormRarityRepository(db *gorm.DB) RarityRepository {
	return &GormRarityRepository{db: db}
}

func (r *GormRarityRepository) Save(rarity *entities.Rarity) error {
	return r.db.Create(&rarity).Error
}

func (r *GormRarityRepository) FindAll() ([]*entities.Rarity, error) {
	var rarityValues []entities.Rarity
	if err := r.db.Find(&rarityValues).Error; err != nil {
		return nil, err
	}

	rarities := make([]*entities.Rarity, len(rarityValues))
	for i := range rarityValues {
		rarities[i] = &rarityValues[i]
	}
	return rarities, nil
}

func (r *GormRarityRepository) FindByID(id string) (*entities.Rarity, error) {
	var rarity entities.Rarity
	if err := r.db.Where("id = ?", id).First(&rarity).Error; err != nil {
		return &entities.Rarity{}, err
	}
	return &rarity, nil
}

func (r *GormRarityRepository) Patch(id string, rarity *entities.Rarity) error {
	if err := r.db.Model(&entities.Rarity{}).Where("id = ?", id).Updates(rarity).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormRarityRepository) Delete(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&entities.Rarity{}).Error; err != nil {
		return err
	}
	return nil
}
