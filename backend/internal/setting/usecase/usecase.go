package usecase

import (
	"encoding/json"
	"time"

	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/MingPV/clean-go-template/internal/setting/repository"
	"github.com/MingPV/clean-go-template/pkg/redisclient"
)

// SettingService
type SettingService struct {
	repo repository.SettingRepository
}

// Init SettingService function
func NewSettingService(repo repository.SettingRepository) SettingUseCase {
	return &SettingService{repo: repo}
}

// SettingService Methods - 1 create
func (s *SettingService) CreateSetting(setting *entities.Setting) error {
	if err := s.repo.Save(setting); err != nil {
		return err
	}

	// Save to Redis cache
	bytes, _ := json.Marshal(setting)
	redisclient.Set("setting:"+setting.ID.String(), string(bytes), time.Minute*10)

	return nil
}

// SettingService Methods - 2 find all
func (s *SettingService) FindAllSettings() ([]*entities.Setting, error) {
	settings, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return settings, nil
}

// SettingService Methods - 3 find by id
func (s *SettingService) FindSettingByID(id string) (*entities.Setting, error) {

	// Check if the setting is in the cache
	jsonData, err := redisclient.Get("setting:" + id)
	if err == nil {
		var setting entities.Setting
		json.Unmarshal([]byte(jsonData), &setting)
		// fmt.Println("Cache hit, returning from cache")
		return &setting, nil
	}

	setting, err := s.repo.FindByID(id)
	if err != nil {
		return &entities.Setting{}, err
	}

	// If not found in the cache, save it to the cache
	// fmt.Println("Cache miss saving to cache")
	bytes, _ := json.Marshal(setting)
	redisclient.Set("setting:"+id, string(bytes), time.Minute*10)

	return setting, nil
}

// SettingService Methods - 4 patch
func (s *SettingService) PatchSetting(id string, setting *entities.Setting) error {

	if err := s.repo.Patch(id, setting); err != nil {
		return err
	}

	// Update cache after patching
	updatedSetting, err := s.repo.FindByID(id)
	if err == nil {
		bytes, _ := json.Marshal(updatedSetting)
		redisclient.Set("setting:"+id, string(bytes), time.Minute*10)
	}

	return nil
}

// SettingService Methods - 5 delete
func (s *SettingService) DeleteSetting(id string) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}

	// Delete cache after removing from DB
	redisclient.Delete("setting:" + id)

	return nil
}
