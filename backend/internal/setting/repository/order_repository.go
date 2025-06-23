package repository

import "github.com/MingPV/clean-go-template/internal/entities"

type SettingRepository interface {
	Save(setting *entities.Setting) error
	FindAll() ([]*entities.Setting, error)
	FindByID(id string) (*entities.Setting, error)
	Patch(id string, setting *entities.Setting) error
	Delete(id string) error
}
