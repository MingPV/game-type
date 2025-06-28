package usecase

import (
	"encoding/json"
	"time"

	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/MingPV/clean-go-template/internal/status/repository"
	"github.com/MingPV/clean-go-template/pkg/redisclient"
)

// StatusService
type StatusService struct {
	repo repository.StatusRepository
}

// Init StatusService function
func NewStatusService(repo repository.StatusRepository) StatusUseCase {
	return &StatusService{repo: repo}
}

// StatusService Methods - 1 create
func (s *StatusService) CreateStatus(status *entities.Status) error {
	if err := s.repo.Save(status); err != nil {
		return err
	}

	// Save to Redis cache
	bytes, _ := json.Marshal(status)
	redisclient.Set("status:"+status.CharacterID.String(), string(bytes), time.Minute*10)

	return nil
}

// StatusService Methods - 2 find all
func (s *StatusService) FindAllStatuses() ([]*entities.Status, error) {
	statuses, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return statuses, nil
}

// StatusService Methods - 3 find by id
func (s *StatusService) FindStatusByCharacterID(character_id string) (*entities.Status, error) {

	// Check if the status is in the cache
	jsonData, err := redisclient.Get("status:" + character_id)
	if err == nil {
		var status entities.Status
		json.Unmarshal([]byte(jsonData), &status)
		// fmt.Println("Cache hit, returning from cache")
		return &status, nil
	}

	status, err := s.repo.FindByCharacterID(character_id)
	if err != nil {
		return &entities.Status{}, err
	}

	// If not found in the cache, save it to the cache
	// fmt.Println("Cache miss saving to cache")
	bytes, _ := json.Marshal(status)
	redisclient.Set("status:"+character_id, string(bytes), time.Minute*10)

	return status, nil
}

// StatusService Methods - 4 patch
func (s *StatusService) PatchStatus(character_id string, status *entities.Status) error {

	if err := s.repo.Patch(character_id, status); err != nil {
		return err
	}

	// Update cache after patching
	updatedStatus, err := s.repo.FindByCharacterID(character_id)
	if err == nil {
		bytes, _ := json.Marshal(updatedStatus)
		redisclient.Set("status:"+character_id, string(bytes), time.Minute*10)
	}

	return nil
}

// StatusService Methods - 5 delete
func (s *StatusService) DeleteStatus(character_id string) error {
	if err := s.repo.Delete(character_id); err != nil {
		return err
	}

	// Delete cache after removing from DB
	redisclient.Delete("status:" + character_id)

	return nil
}
