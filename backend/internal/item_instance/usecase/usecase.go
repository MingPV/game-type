package usecase

import (
	"encoding/json"
	"time"

	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/MingPV/clean-go-template/internal/item_instance/repository"
	"github.com/MingPV/clean-go-template/pkg/redisclient"
)

// ItemInstanceService
type ItemInstanceService struct {
	repo repository.ItemInstanceRepository
}

// Init ItemInstanceService function
func NewItemInstanceService(repo repository.ItemInstanceRepository) ItemInstanceUseCase {
	return &ItemInstanceService{repo: repo}
}

// ItemInstanceService Methods - 1 create
func (s *ItemInstanceService) CreateItemInstance(itemInstance *entities.ItemInstance) error {
	if err := s.repo.Save(itemInstance); err != nil {
		return err
	}

	// Save to Redis cache
	bytes, _ := json.Marshal(itemInstance)
	redisclient.Set("itemInstance:"+itemInstance.ID.String(), string(bytes), time.Minute*10)

	return nil
}

// ItemInstanceService Methods - 2 find all
func (s *ItemInstanceService) FindAllItemInstances() ([]*entities.ItemInstance, error) {
	itemInstances, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return itemInstances, nil
}

// ItemInstanceService Methods - 3 find by id
func (s *ItemInstanceService) FindItemInstanceByID(id string) (*entities.ItemInstance, error) {

	// Check if the itemInstance is in the cache
	jsonData, err := redisclient.Get("itemInstance:" + id)
	if err == nil {
		var itemInstance entities.ItemInstance
		json.Unmarshal([]byte(jsonData), &itemInstance)
		// fmt.Println("Cache hit, returning from cache")
		return &itemInstance, nil
	}

	itemInstance, err := s.repo.FindByID(id)
	if err != nil {
		return &entities.ItemInstance{}, err
	}

	// If not found in the cache, save it to the cache
	// fmt.Println("Cache miss saving to cache")
	bytes, _ := json.Marshal(itemInstance)
	redisclient.Set("itemInstance:"+id, string(bytes), time.Minute*10)

	return itemInstance, nil
}

// ItemInstanceService Methods - 4 patch
func (s *ItemInstanceService) PatchItemInstance(id string, itemInstance *entities.ItemInstance) error {

	if err := s.repo.Patch(id, itemInstance); err != nil {
		return err
	}

	// Update cache after patching
	updatedItemInstance, err := s.repo.FindByID(id)
	if err == nil {
		bytes, _ := json.Marshal(updatedItemInstance)
		redisclient.Set("itemInstance:"+id, string(bytes), time.Minute*10)
	}

	return nil
}

// ItemInstanceService Methods - 5 delete
func (s *ItemInstanceService) DeleteItemInstance(id string) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}

	// Delete cache after removing from DB
	redisclient.Delete("itemInstance:" + id)

	return nil
}
