package usecase

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/MingPV/clean-go-template/internal/item_type/repository"
	"github.com/MingPV/clean-go-template/pkg/redisclient"
)

// ItemTypeService
type ItemTypeService struct {
	repo repository.ItemTypeRepository
}

// Init ItemTypeService function
func NewItemTypeService(repo repository.ItemTypeRepository) ItemTypeUseCase {
	return &ItemTypeService{repo: repo}
}

// ItemTypeService Methods - 1 create
func (s *ItemTypeService) CreateItemType(itemType *entities.ItemType) error {
	if err := s.repo.Save(itemType); err != nil {
		return err
	}

	// Save to Redis cache
	bytes, _ := json.Marshal(itemType)
	redisclient.Set("itemType:"+itemType.ID.String(), string(bytes), time.Minute*10)

	return nil
}

// ItemTypeService Methods - 2 find all
func (s *ItemTypeService) FindAllItemTypes() ([]*entities.ItemType, error) {
	itemTypes, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return itemTypes, nil
}

// ItemTypeService Methods - 3 find by id
func (s *ItemTypeService) FindItemTypeByID(id int) (*entities.ItemType, error) {

	// Check if the itemType is in the cache
	jsonData, err := redisclient.Get("itemType:" + strconv.Itoa(id))
	if err == nil {
		var itemType entities.ItemType
		json.Unmarshal([]byte(jsonData), &itemType)
		// fmt.Println("Cache hit, returning from cache")
		return &itemType, nil
	}

	itemType, err := s.repo.FindByID(id)
	if err != nil {
		return &entities.ItemType{}, err
	}

	// If not found in the cache, save it to the cache
	// fmt.Println("Cache miss saving to cache")
	bytes, _ := json.Marshal(itemType)
	redisclient.Set("itemType:"+strconv.Itoa(id), string(bytes), time.Minute*10)

	return itemType, nil
}

// ItemTypeService Methods - 4 patch
// func (s *ItemTypeService) PatchItemType(id int, itemType *entities.ItemType) error {
// 	if itemType.Total <= 0 {
// 		return errors.New("total must be positive")
// 	}
// 	if err := s.repo.Patch(id, itemType); err != nil {
// 		return err
// 	}

// 	// Update cache after patching
// 	updatedItemType, err := s.repo.FindByID(id)
// 	if err == nil {
// 		bytes, _ := json.Marshal(updatedItemType)
// 		redisclient.Set("itemType:"+strconv.Itoa(id), string(bytes), time.Minute*10)
// 	}

// 	return nil
// }

// ItemTypeService Methods - 5 delete
func (s *ItemTypeService) DeleteItemType(id int) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}

	// Delete cache after removing from DB
	redisclient.Delete("itemType:" + strconv.Itoa(id))

	return nil
}
