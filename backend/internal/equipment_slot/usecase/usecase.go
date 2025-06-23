package usecase

import (
	"encoding/json"
	"time"

	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/MingPV/clean-go-template/internal/equipment_slot/repository"
	"github.com/MingPV/clean-go-template/pkg/redisclient"
)

// EquipmentSlotService
type EquipmentSlotService struct {
	repo repository.EquipmentSlotRepository
}

// Init EquipmentSlotService function
func NewEquipmentSlotService(repo repository.EquipmentSlotRepository) EquipmentSlotUseCase {
	return &EquipmentSlotService{repo: repo}
}

// EquipmentSlotService Methods - 1 create
func (s *EquipmentSlotService) CreateEquipmentSlot(equipmentSlot *entities.EquipmentSlot) error {
	if err := s.repo.Save(equipmentSlot); err != nil {
		return err
	}

	// Save to Redis cache
	bytes, _ := json.Marshal(equipmentSlot)
	redisclient.Set("equipmentSlot:"+equipmentSlot.ID.String(), string(bytes), time.Minute*10)

	return nil
}

// EquipmentSlotService Methods - 2 find all
func (s *EquipmentSlotService) FindAllEquipmentSlots() ([]*entities.EquipmentSlot, error) {
	equipmentSlots, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return equipmentSlots, nil
}

// EquipmentSlotService Methods - 3 find by id
func (s *EquipmentSlotService) FindEquipmentSlotByID(id string) (*entities.EquipmentSlot, error) {

	// Check if the equipmentSlot is in the cache
	jsonData, err := redisclient.Get("equipmentSlot:" + id)
	if err == nil {
		var equipmentSlot entities.EquipmentSlot
		json.Unmarshal([]byte(jsonData), &equipmentSlot)
		// fmt.Println("Cache hit, returning from cache")
		return &equipmentSlot, nil
	}

	equipmentSlot, err := s.repo.FindByID(id)
	if err != nil {
		return &entities.EquipmentSlot{}, err
	}

	// If not found in the cache, save it to the cache
	// fmt.Println("Cache miss saving to cache")
	bytes, _ := json.Marshal(equipmentSlot)
	redisclient.Set("equipmentSlot:"+id, string(bytes), time.Minute*10)

	return equipmentSlot, nil
}

// EquipmentSlotService Methods - 4 patch
// func (s *EquipmentSlotService) PatchEquipmentSlot(id int, equipmentSlot *entities.EquipmentSlot) error {
// 	if equipmentSlot.Total <= 0 {
// 		return errors.New("total must be positive")
// 	}
// 	if err := s.repo.Patch(id, equipmentSlot); err != nil {
// 		return err
// 	}

// 	// Update cache after patching
// 	updatedEquipmentSlot, err := s.repo.FindByID(id)
// 	if err == nil {
// 		bytes, _ := json.Marshal(updatedEquipmentSlot)
// 		redisclient.Set("equipmentSlot:"+strconv.Itoa(id), string(bytes), time.Minute*10)
// 	}

// 	return nil
// }

// EquipmentSlotService Methods - 5 delete
func (s *EquipmentSlotService) DeleteEquipmentSlot(id string) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}

	// Delete cache after removing from DB
	redisclient.Delete("equipmentSlot:" + id)

	return nil
}
