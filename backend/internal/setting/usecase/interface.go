package usecase

import "github.com/MingPV/clean-go-template/internal/entities"

type SettingUseCase interface {
	FindAllSettings() ([]*entities.Setting, error)
	CreateSetting(setting *entities.Setting) error
	PatchSetting(id string, setting *entities.Setting) error
	DeleteSetting(id string) error
	FindSettingByID(id string) (*entities.Setting, error)
}
