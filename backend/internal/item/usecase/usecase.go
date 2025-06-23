package usecase

import (
	"encoding/json"
	"time"

	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/MingPV/clean-go-template/internal/item/repository"
	"github.com/MingPV/clean-go-template/pkg/redisclient"
)

// ItemService
type ItemService struct {
	repo repository.ItemRepository
}

// Init ItemService function
func NewItemService(repo repository.ItemRepository) ItemUseCase {
	return &ItemService{repo: repo}
}

// ItemService Methods - 1 create
func (s *ItemService) CreateItem(item *entities.Item) error {
	if err := s.repo.Save(item); err != nil {
		return err
	}

	// Save to Redis cache
	bytes, _ := json.Marshal(item)
	redisclient.Set("item:"+item.ID.String(), string(bytes), time.Minute*10)

	return nil
}

// ItemService Methods - 2 find all
func (s *ItemService) FindAllItems() ([]*entities.Item, error) {
	items, err := s.repo.FindAll()
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

	item, err := s.repo.FindByID(id)
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

	if err := s.repo.Patch(id, item); err != nil {
		return err
	}

	// Update cache after patching
	updatedItem, err := s.repo.FindByID(id)
	if err == nil {
		bytes, _ := json.Marshal(updatedItem)
		redisclient.Set("item:"+id, string(bytes), time.Minute*10)
	}

	return nil
}

// ItemService Methods - 5 delete
func (s *ItemService) DeleteItem(id string) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}

	// Delete cache after removing from DB
	redisclient.Delete("item:" + id)

	return nil
}
