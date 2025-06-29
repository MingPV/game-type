package usecase

import (
	"encoding/json"
	"time"

	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/MingPV/clean-go-template/internal/monster/repository"
	"github.com/MingPV/clean-go-template/pkg/redisclient"
)

// MonsterService
type MonsterService struct {
	repo repository.MonsterRepository
}

// Init MonsterService function
func NewMonsterService(repo repository.MonsterRepository) MonsterUseCase {
	return &MonsterService{repo: repo}
}

// MonsterService Methods - 1 create
func (s *MonsterService) CreateMonster(monster *entities.Monster) (*entities.Monster, error) {
	if err := s.repo.Save(monster); err != nil {
		return nil, err
	}

	monster_return, err := s.repo.FindByID(monster.ID.String())

	if err != nil {
		return nil, err
	}

	// Save to Redis cache
	bytes, _ := json.Marshal(monster_return)
	redisclient.Set("monster:"+monster_return.ID.String(), string(bytes), time.Minute*10)

	return monster_return, nil
}

// MonsterService Methods - 2 find all
func (s *MonsterService) FindAllMonsters() ([]*entities.Monster, error) {
	monsters, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return monsters, nil
}

// MonsterService Methods - 3 find by id
func (s *MonsterService) FindMonsterByID(id string) (*entities.Monster, error) {

	// Check if the monster is in the cache
	jsonData, err := redisclient.Get("monster:" + id)
	if err == nil {
		var monster entities.Monster
		json.Unmarshal([]byte(jsonData), &monster)
		// fmt.Println("Cache hit, returning from cache")
		return &monster, nil
	}

	monster, err := s.repo.FindByID(id)
	if err != nil {
		return &entities.Monster{}, err
	}

	// If not found in the cache, save it to the cache
	// fmt.Println("Cache miss saving to cache")
	bytes, _ := json.Marshal(monster)
	redisclient.Set("monster:"+id, string(bytes), time.Minute*10)

	return monster, nil
}

// MonsterService Methods - 4 patch
func (s *MonsterService) PatchMonster(id string, monster *entities.Monster) error {

	if err := s.repo.Patch(id, monster); err != nil {
		return err
	}

	// Update cache after patching
	updatedMonster, err := s.repo.FindByID(id)
	if err == nil {
		bytes, _ := json.Marshal(updatedMonster)
		redisclient.Set("monster:"+id, string(bytes), time.Minute*10)
	}

	return nil
}

// MonsterService Methods - 5 delete
func (s *MonsterService) DeleteMonster(id string) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}

	// Delete cache after removing from DB
	redisclient.Delete("monster:" + id)

	return nil
}
