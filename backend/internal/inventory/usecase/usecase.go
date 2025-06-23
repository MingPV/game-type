package usecase

import (
	"encoding/json"
	"time"

	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/MingPV/clean-go-template/internal/inventory/repository"
	"github.com/MingPV/clean-go-template/pkg/redisclient"
)

// InventoryService
type InventoryService struct {
	repo repository.InventoryRepository
}

// Init InventoryService function
func NewInventoryService(repo repository.InventoryRepository) InventoryUseCase {
	return &InventoryService{repo: repo}
}

// InventoryService Methods - 1 create
func (s *InventoryService) CreateInventory(inventory *entities.Inventory) error {
	if err := s.repo.Save(inventory); err != nil {
		return err
	}

	// Save to Redis cache
	bytes, _ := json.Marshal(inventory)
	redisclient.Set("inventory:"+inventory.ID.String(), string(bytes), time.Minute*10)

	return nil
}

// InventoryService Methods - 2 find all
func (s *InventoryService) FindAllInventories() ([]*entities.Inventory, error) {
	inventories, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return inventories, nil
}

// InventoryService Methods - 3 find by id
func (s *InventoryService) FindInventoryByID(id string) (*entities.Inventory, error) {

	// Check if the inventory is in the cache
	jsonData, err := redisclient.Get("inventory:" + id)
	if err == nil {
		var inventory entities.Inventory
		json.Unmarshal([]byte(jsonData), &inventory)
		// fmt.Println("Cache hit, returning from cache")
		return &inventory, nil
	}

	inventory, err := s.repo.FindByID(id)
	if err != nil {
		return &entities.Inventory{}, err
	}

	// If not found in the cache, save it to the cache
	// fmt.Println("Cache miss saving to cache")
	bytes, _ := json.Marshal(inventory)
	redisclient.Set("inventory:"+id, string(bytes), time.Minute*10)

	return inventory, nil
}

// InventoryService Methods - 4 patch
func (s *InventoryService) PatchInventory(id string, inventory *entities.Inventory) error {

	if err := s.repo.Patch(id, inventory); err != nil {
		return err
	}

	// Update cache after patching
	updatedInventory, err := s.repo.FindByID(id)
	if err == nil {
		bytes, _ := json.Marshal(updatedInventory)
		redisclient.Set("inventory:"+id, string(bytes), time.Minute*10)
	}

	return nil
}

// InventoryService Methods - 5 delete
func (s *InventoryService) DeleteInventory(id string) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}

	// Delete cache after removing from DB
	redisclient.Delete("inventory:" + id)

	return nil
}
