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
	if err := r.db.Preload("Class").Preload("Status").Find(&characterValues).Error; err != nil {
		return nil, err
	}

	characters := make([]*entities.Character, len(characterValues))
	for i := range characterValues {
		characters[i] = &characterValues[i]
	}
	return characters, nil
}

func (r *GormCharacterRepository) FindByID(id int) (*entities.Character, error) {
	var character entities.Character
	if err := r.db.First(&character, id).Error; err != nil {
		return &entities.Character{}, err
	}
	return &character, nil
}

func (r *GormCharacterRepository) Patch(id int, character *entities.Character) error {
	if err := r.db.Model(&entities.Character{}).Where("id = ?", id).Updates(character).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormCharacterRepository) Delete(id int) error {
	if err := r.db.Delete(&entities.Character{}, id).Error; err != nil {
		return err
	}
	return nil
}
