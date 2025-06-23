package usecase

import (
	"encoding/json"
	"time"

	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/MingPV/clean-go-template/internal/monster_type/repository"
	"github.com/MingPV/clean-go-template/pkg/redisclient"
)

// MonsterTypeService
type MonsterTypeService struct {
	repo repository.MonsterTypeRepository
}

// Init MonsterTypeService function
func NewMonsterTypeService(repo repository.MonsterTypeRepository) MonsterTypeUseCase {
	return &MonsterTypeService{repo: repo}
}

// MonsterTypeService Methods - 1 create
func (s *MonsterTypeService) CreateMonsterType(monsterType *entities.MonsterType) error {
	if err := s.repo.Save(monsterType); err != nil {
		return err
	}

	// Save to Redis cache
	bytes, _ := json.Marshal(monsterType)
	redisclient.Set("monsterType:"+monsterType.ID.String(), string(bytes), time.Minute*10)

	return nil
}

// MonsterTypeService Methods - 2 find all
func (s *MonsterTypeService) FindAllMonsterTypes() ([]*entities.MonsterType, error) {
	monsterTypes, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return monsterTypes, nil
}

// MonsterTypeService Methods - 3 find by id
func (s *MonsterTypeService) FindMonsterTypeByID(id string) (*entities.MonsterType, error) {

	// Check if the monsterType is in the cache
	jsonData, err := redisclient.Get("monsterType:" + id)
	if err == nil {
		var monsterType entities.MonsterType
		json.Unmarshal([]byte(jsonData), &monsterType)
		// fmt.Println("Cache hit, returning from cache")
		return &monsterType, nil
	}

	monsterType, err := s.repo.FindByID(id)
	if err != nil {
		return &entities.MonsterType{}, err
	}

	// If not found in the cache, save it to the cache
	// fmt.Println("Cache miss saving to cache")
	bytes, _ := json.Marshal(monsterType)
	redisclient.Set("monsterType:"+id, string(bytes), time.Minute*10)

	return monsterType, nil
}

// MonsterTypeService Methods - 4 patch
func (s *MonsterTypeService) PatchMonsterType(id string, monsterType *entities.MonsterType) error {

	if err := s.repo.Patch(id, monsterType); err != nil {
		return err
	}

	// Update cache after patching
	updatedMonsterType, err := s.repo.FindByID(id)
	if err == nil {
		bytes, _ := json.Marshal(updatedMonsterType)
		redisclient.Set("monsterType:"+id, string(bytes), time.Minute*10)
	}

	return nil
}

// MonsterTypeService Methods - 5 delete
func (s *MonsterTypeService) DeleteMonsterType(id string) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}

	// Delete cache after removing from DB
	redisclient.Delete("monsterType:" + id)

	return nil
}
