package usecase

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/MingPV/clean-go-template/internal/entities"
	itemRepo "github.com/MingPV/clean-go-template/internal/item/repository"
	itemLevelStatRepo "github.com/MingPV/clean-go-template/internal/item_level_stat/repository"
	"github.com/MingPV/clean-go-template/pkg/redisclient"
)

// ItemService
type ItemService struct {
	itemReposiroty          itemRepo.ItemRepository
	itemLevelStatRepository itemLevelStatRepo.ItemLevelStatRepository
}

// Init ItemService function
func NewItemService(item_repo itemRepo.ItemRepository, item_level_stat_repo itemLevelStatRepo.ItemLevelStatRepository) ItemUseCase {
	return &ItemService{itemReposiroty: item_repo, itemLevelStatRepository: item_level_stat_repo}
}

// ItemService Methods - 1 create
func (s *ItemService) CreateItem(item *entities.Item, item_level_stat *entities.ItemLevelStat) (*entities.Item, error) {

	if err := s.itemLevelStatRepository.Save(item_level_stat); err != nil {
		return nil, err
	}

	if err := s.itemReposiroty.Save(item); err != nil {
		return nil, err
	}

	item_return, err := s.itemReposiroty.FindByID(item.ID.String())

	if err != nil {
		return nil, err
	}

	// Save to Redis cache
	bytes, _ := json.Marshal(item_return)
	redisclient.Set("item:"+item.ID.String(), string(bytes), time.Minute*10)

	return item_return, nil
}

// ItemService Methods - 2 find all
func (s *ItemService) FindAllItems() ([]*entities.Item, error) {
	items, err := s.itemReposiroty.FindAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

// ItemService Methods - 3 find by id
func (s *ItemService) FindItemByID(id string) (*entities.Item, error) {

	// Check if the item is in the cache
	jsonData, err := redisclient.Get("item:" + id)
	if err == nil {
		var item entities.Item
		json.Unmarshal([]byte(jsonData), &item)
		// fmt.Println("Cache hit, returning from cache")
		return &item, nil
	}

	item, err := s.itemReposiroty.FindByID(id)
	if err != nil {
		return &entities.Item{}, err
	}

	// If not found in the cache, save it to the cache
	// fmt.Println("Cache miss saving to cache")
	bytes, _ := json.Marshal(item)
	redisclient.Set("item:"+id, string(bytes), time.Minute*10)

	return item, nil
}

// ItemService Methods - 4 patch
func (s *ItemService) PatchItem(id string, item *entities.Item) error {

	fmt.Println("item patched2 :: ", item)

	if err := s.itemReposiroty.Patch(id, item); err != nil {
		return err
	}

	fmt.Println("item patched :: ", item)

	// Update cache after patching
	updatedItem, err := s.itemReposiroty.FindByID(id)
	if err == nil {
		bytes, _ := json.Marshal(updatedItem)
		redisclient.Set("item:"+id, string(bytes), time.Minute*10)
	}

	return nil
}

// ItemService Methods - 5 delete
func (s *ItemService) DeleteItem(id string) error {
	if err := s.itemReposiroty.Delete(id); err != nil {
		return err
	}

	// Delete cache after removing from DB
	redisclient.Delete("item:" + id)

	return nil
}
