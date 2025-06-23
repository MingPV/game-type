package repository

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	"gorm.io/gorm"
)

type GormClassRepository struct {
	db *gorm.DB
}

func NewGormClassRepository(db *gorm.DB) ClassRepository {
	return &GormClassRepository{db: db}
}

func (r *GormClassRepository) Save(class *entities.Class) error {
	return r.db.Create(&class).Error
}

func (r *GormClassRepository) FindAll() ([]*entities.Class, error) {
	var classValues []entities.Class
	if err := r.db.Find(&classValues).Error; err != nil {
		return nil, err
	}

	classes := make([]*entities.Class, len(classValues))
	for i := range classValues {
		classes[i] = &classValues[i]
	}
	return classes, nil
}

func (r *GormClassRepository) FindByID(id string) (*entities.Class, error) {
	var class entities.Class
	if err := r.db.Where("id = ?", id).First(&class).Error; err != nil {
		return &entities.Class{}, err
	}
	return &class, nil
}

func (r *GormClassRepository) Patch(id int, class *entities.Class) error {
	if err := r.db.Model(&entities.Class{}).Where("id = ?", id).Updates(class).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormClassRepository) Delete(id int) error {
	if err := r.db.Delete(&entities.Class{}, id).Error; err != nil {
		return err
	}
	return nil
}
