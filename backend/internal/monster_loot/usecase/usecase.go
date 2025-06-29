package usecase

import (
	"encoding/json"
	"time"

	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/MingPV/clean-go-template/internal/monster_loot/repository"
	"github.com/MingPV/clean-go-template/pkg/redisclient"
)

// MonsterLootService
type MonsterLootService struct {
	repo repository.MonsterLootRepository
}

// Init MonsterLootService function
func NewMonsterLootService(repo repository.MonsterLootRepository) MonsterLootUseCase {
	return &MonsterLootService{repo: repo}
}

// MonsterLootService Methods - 1 create
func (s *MonsterLootService) CreateMonsterLoot(monsterLoot *entities.MonsterLoot) (*entities.MonsterLoot, error) {
	if err := s.repo.Save(monsterLoot); err != nil {
		return nil, err
	}

	loot_return, err := s.repo.FindByMonsterIDAndItemID(monsterLoot.MonsterID.String(), monsterLoot.ItemID.String())

	if err != nil {
		return nil, err
	}

	// Save to Redis cache
	bytes, _ := json.Marshal(loot_return)
	redisclient.Set("monsterLoot:"+monsterLoot.MonsterID.String()+monsterLoot.ItemID.String(), string(bytes), time.Minute*10)

	return loot_return, nil
}

// MonsterLootService Methods - 2 find all
func (s *MonsterLootService) FindAllMonsterLoots() ([]*entities.MonsterLoot, error) {
	monsterLoots, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return monsterLoots, nil
}

// Need to use 2 primary key now it's not work!!
// MonsterLootService Methods - 3 find by id
// func (s *MonsterLootService) FindMonsterLootByID(id string) (*entities.MonsterLoot, error) {

// 	// Check if the monsterLoot is in the cache
// 	jsonData, err := redisclient.Get("monsterLoot:" + id)
// 	if err == nil {
// 		var monsterLoot entities.MonsterLoot
// 		json.Unmarshal([]byte(jsonData), &monsterLoot)
// 		// fmt.Println("Cache hit, returning from cache")
// 		return &monsterLoot, nil
// 	}

// 	monsterLoot, err := s.repo.FindByID(id)
// 	if err != nil {
// 		return &entities.MonsterLoot{}, err
// 	}

// 	// If not found in the cache, save it to the cache
// 	// fmt.Println("Cache miss saving to cache")
// 	bytes, _ := json.Marshal(monsterLoot)
// 	redisclient.Set("monsterLoot:"+id, string(bytes), time.Minute*10)

// 	return monsterLoot, nil
// }

func (s *MonsterLootService) FindMonsterLootByMonsterID(monster_id string) ([]*entities.MonsterLoot, error) {

	// Check if the monsterLoot is in the cache
	jsonData, err := redisclient.Get("monsterLoot:" + monster_id)
	if err == nil {
		var monsterLoots []*entities.MonsterLoot
		json.Unmarshal([]byte(jsonData), &monsterLoots)
		// fmt.Println("Cache hit, returning from cache")
		return monsterLoots, nil
	}

	monsterLoots, err := s.repo.FindByMonsterID(monster_id)
	if err != nil {
		return nil, err
	}

	// If not found in the cache, save it to the cache
	// fmt.Println("Cache miss saving to cache")
	bytes, _ := json.Marshal(monsterLoots)
	redisclient.Set("monsterLoot:"+monster_id, string(bytes), time.Minute*10)

	return monsterLoots, nil
}

func (s *MonsterLootService) FindMonsterLootByItemID(item_id string) ([]*entities.MonsterLoot, error) {

	// Check if the monsterLoot is in the cache
	jsonData, err := redisclient.Get("monsterLoot:" + item_id)
	if err == nil {
		var monsterLoots []*entities.MonsterLoot
		json.Unmarshal([]byte(jsonData), &monsterLoots)
		// fmt.Println("Cache hit, returning from cache")
		return monsterLoots, nil
	}

	monsterLoots, err := s.repo.FindByItemID(item_id)
	if err != nil {
		return nil, err
	}

	// If not found in the cache, save it to the cache
	// fmt.Println("Cache miss saving to cache")
	bytes, _ := json.Marshal(monsterLoots)
	redisclient.Set("monsterLoot:"+item_id, string(bytes), time.Minute*10)

	return monsterLoots, nil
}

func (s *MonsterLootService) FindMonsterLootByMonsterIDAndItemID(monsterID, itemID string) (*entities.MonsterLoot, error) {
	cacheKey := "monsterLoot:" + monsterID + ":" + itemID

	// Check cache first
	jsonData, err := redisclient.Get(cacheKey)
	if err == nil {
		var monsterLoot entities.MonsterLoot
		if err := json.Unmarshal([]byte(jsonData), &monsterLoot); err == nil {
			return &monsterLoot, nil
		}
	}

	// Fetch from repository
	monsterLoot, err := s.repo.FindByMonsterIDAndItemID(monsterID, itemID)
	if err != nil {
		return nil, err
	}

	// Cache the result
	bytes, _ := json.Marshal(monsterLoot)
	redisclient.Set(cacheKey, string(bytes), time.Minute*10)

	return monsterLoot, nil
}

// MonsterLootService Methods - 4 patch
func (s *MonsterLootService) PatchMonsterLoot(monster_id string, item_id string, monsterLoot *entities.MonsterLoot) error {
	// Patch the monster loot using both primary keys
	if err := s.repo.Patch(monster_id, item_id, monsterLoot); err != nil {
		return err
	}

	// Update cache after patching
	updatedMonsterLoot, err := s.repo.FindByMonsterIDAndItemID(monster_id, item_id)
	if err == nil {
		cacheKey := "monsterLoot:" + monster_id + ":" + item_id
		bytes, _ := json.Marshal(updatedMonsterLoot)
		redisclient.Set(cacheKey, string(bytes), time.Minute*10)
	}

	return nil
}

// MonsterLootService Methods - 5 delete
func (s *MonsterLootService) DeleteMonsterLoot(monster_id string, item_id string) error {
	// Delete from repository using both primary keys
	if err := s.repo.Delete(monster_id, item_id); err != nil {
		return err
	}

	// Delete cache after removing from DB
	cacheKey := "monsterLoot:" + monster_id + ":" + item_id
	redisclient.Delete(cacheKey)

	return nil
}
