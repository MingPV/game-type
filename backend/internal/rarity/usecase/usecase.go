package usecase

import (
	"encoding/json"
	"time"

	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/MingPV/clean-go-template/internal/rarity/repository"
	"github.com/MingPV/clean-go-template/pkg/redisclient"
)

// RarityService
type RarityService struct {
	repo repository.RarityRepository
}

// Init RarityService function
func NewRarityService(repo repository.RarityRepository) RarityUseCase {
	return &RarityService{repo: repo}
}

// RarityService Methods - 1 create
func (s *RarityService) CreateRarity(rarity *entities.Rarity) error {
	if err := s.repo.Save(rarity); err != nil {
		return err
	}

	// Save to Redis cache
	bytes, _ := json.Marshal(rarity)
	redisclient.Set("rarity:"+rarity.ID.String(), string(bytes), time.Minute*10)

	return nil
}

// RarityService Methods - 2 find all
func (s *RarityService) FindAllRarities() ([]*entities.Rarity, error) {
	rarities, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return rarities, nil
}

// RarityService Methods - 3 find by id
func (s *RarityService) FindRarityByID(id string) (*entities.Rarity, error) {

	// Check if the rarity is in the cache
	jsonData, err := redisclient.Get("rarity:" + id)
	if err == nil {
		var rarity entities.Rarity
		json.Unmarshal([]byte(jsonData), &rarity)
		// fmt.Println("Cache hit, returning from cache")
		return &rarity, nil
	}

	rarity, err := s.repo.FindByID(id)
	if err != nil {
		return &entities.Rarity{}, err
	}

	// If not found in the cache, save it to the cache
	// fmt.Println("Cache miss saving to cache")
	bytes, _ := json.Marshal(rarity)
	redisclient.Set("rarity:"+id, string(bytes), time.Minute*10)

	return rarity, nil
}

// RarityService Methods - 4 patch
func (s *RarityService) PatchRarity(id string, rarity *entities.Rarity) error {

	if err := s.repo.Patch(id, rarity); err != nil {
		return err
	}

	// Update cache after patching
	updatedRarity, err := s.repo.FindByID(id)
	if err == nil {
		bytes, _ := json.Marshal(updatedRarity)
		redisclient.Set("rarity:"+id, string(bytes), time.Minute*10)
	}

	return nil
}

// RarityService Methods - 5 delete
func (s *RarityService) DeleteRarity(id string) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}

	// Delete cache after removing from DB
	redisclient.Delete("rarity:" + id)

	return nil
}
