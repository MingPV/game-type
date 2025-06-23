package usecase

import (
	"encoding/json"
	"time"

	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/MingPV/clean-go-template/internal/item_level_stat/repository"
	"github.com/MingPV/clean-go-template/pkg/redisclient"
)

// ItemLevelStatService
type ItemLevelStatService struct {
	repo repository.ItemLevelStatRepository
}

// Init ItemLevelStatService function
func NewItemLevelStatService(repo repository.ItemLevelStatRepository) ItemLevelStatUseCase {
	return &ItemLevelStatService{repo: repo}
}

// ItemLevelStatService Methods - 1 create
func (s *ItemLevelStatService) CreateItemLevelStat(itemLevelStat *entities.ItemLevelStat) error {
	if err := s.repo.Save(itemLevelStat); err != nil {
		return err
	}

	// Save to Redis cache
	bytes, _ := json.Marshal(itemLevelStat)
	redisclient.Set("itemLevelStat:"+itemLevelStat.ItemID.String(), string(bytes), time.Minute*10)

	return nil
}

// ItemLevelStatService Methods - 2 find all
func (s *ItemLevelStatService) FindAllItemLevelStats() ([]*entities.ItemLevelStat, error) {
	itemLevelStats, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return itemLevelStats, nil
}

// ItemLevelStatService Methods - 3 find by id
func (s *ItemLevelStatService) FindItemLevelStatByID(id string) (*entities.ItemLevelStat, error) {

	// Check if the itemLevelStat is in the cache
	jsonData, err := redisclient.Get("itemLevelStat:" + id)
	if err == nil {
		var itemLevelStat entities.ItemLevelStat
		json.Unmarshal([]byte(jsonData), &itemLevelStat)
		// fmt.Println("Cache hit, returning from cache")
		return &itemLevelStat, nil
	}

	itemLevelStat, err := s.repo.FindByID(id)
	if err != nil {
		return &entities.ItemLevelStat{}, err
	}

	// If not found in the cache, save it to the cache
	// fmt.Println("Cache miss saving to cache")
	bytes, _ := json.Marshal(itemLevelStat)
	redisclient.Set("itemLevelStat:"+id, string(bytes), time.Minute*10)

	return itemLevelStat, nil
}

// ItemLevelStatService Methods - 4 patch
func (s *ItemLevelStatService) PatchItemLevelStat(id string, itemLevelStat *entities.ItemLevelStat) error {

	if err := s.repo.Patch(id, itemLevelStat); err != nil {
		return err
	}

	// Update cache after patching
	updatedItemLevelStat, err := s.repo.FindByID(id)
	if err == nil {
		bytes, _ := json.Marshal(updatedItemLevelStat)
		redisclient.Set("itemLevelStat:"+id, string(bytes), time.Minute*10)
	}

	return nil
}

// ItemLevelStatService Methods - 5 delete
func (s *ItemLevelStatService) DeleteItemLevelStat(id string) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}

	// Delete cache after removing from DB
	redisclient.Delete("itemLevelStat:" + id)

	return nil
}
