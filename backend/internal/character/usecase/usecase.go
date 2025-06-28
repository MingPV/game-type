package usecase

import (
	"encoding/json"
	"time"

	characterRepo "github.com/MingPV/clean-go-template/internal/character/repository"
	"github.com/MingPV/clean-go-template/internal/constants"
	"github.com/MingPV/clean-go-template/internal/entities"
	statusRepo "github.com/MingPV/clean-go-template/internal/status/repository"

	equipmentSlotRepo "github.com/MingPV/clean-go-template/internal/equipment_slot/repository"
	inventoryRepo "github.com/MingPV/clean-go-template/internal/inventory/repository"
	"github.com/MingPV/clean-go-template/pkg/redisclient"
)

// CharacterService
type CharacterService struct {
	characterRepository characterRepo.CharacterRepository
	statusRepository    statusRepo.StatusRepository
	inventoryRepository inventoryRepo.InventoryRepository
	equipmentRepository equipmentSlotRepo.EquipmentSlotRepository
}

// Init CharacterService function
func NewCharacterService(character_repo characterRepo.CharacterRepository, status_repo statusRepo.StatusRepository, inventory_repo inventoryRepo.InventoryRepository, equipment_slot equipmentSlotRepo.EquipmentSlotRepository) CharacterUseCase {
	return &CharacterService{
		characterRepository: character_repo,
		statusRepository:    status_repo,
		equipmentRepository: equipment_slot,
		inventoryRepository: inventory_repo,
	}
}

// CharacterService Methods - 1 create
func (s *CharacterService) CreateCharacter(character *entities.Character) (*entities.Character, error) {

	baseStatus := &entities.Status{
		CharacterID: character.ID,
		StatusPoint: constants.BASE_STATUS_POINTS,
		STR:         constants.STR,
		AGI:         constants.AGI,
		INT:         constants.INT,
		DEX:         constants.DEX,
		VIT:         constants.VIT,
		LUK:         constants.LUK,
	}

	inventory := &entities.Inventory{
		MaxSlots: constants.BAS_MAX_INVENTORY_SLOTS,
	}

	if err := s.statusRepository.Save(baseStatus); err != nil {
		return nil, err
	}

	if err := s.inventoryRepository.Save(inventory); err != nil {
		return nil, err
	}

	// Need to insert inventory first to get the ID
	character.InventoryID = inventory.ID

	if err := s.characterRepository.Save(character); err != nil {
		return nil, err
	}

	character_return, err := s.characterRepository.FindByID(character.ID.String())

	// Save to Redis cache
	bytes, _ := json.Marshal(character_return)
	redisclient.Set("character:"+character.ID.String(), string(bytes), time.Minute*10)

	return character_return, err
}

// CharacterService Methods - 2 find all
func (s *CharacterService) FindAllCharacters() ([]*entities.Character, error) {
	characters, err := s.characterRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return characters, nil
}

// CharacterService Methods - 3 find by id
func (s *CharacterService) FindCharacterByID(id string) (*entities.Character, error) {

	// Check if the character is in the cache
	jsonData, err := redisclient.Get("character:" + id)
	if err == nil {
		var character entities.Character
		json.Unmarshal([]byte(jsonData), &character)
		// fmt.Println("Cache hit, returning from cache")
		return &character, nil
	}

	character, err := s.characterRepository.FindByID(id)
	if err != nil {
		return &entities.Character{}, err
	}

	// If not found in the cache, save it to the cache
	// fmt.Println("Cache miss saving to cache")
	bytes, _ := json.Marshal(character)
	redisclient.Set("character:"+id, string(bytes), time.Minute*10)

	return character, nil
}

// CharacterService Methods - 4 patch
func (s *CharacterService) PatchCharacter(id string, character *entities.Character) error {

	if err := s.characterRepository.Patch(id, character); err != nil {
		return err
	}

	// Update cache after patching
	updatedCharacter, err := s.characterRepository.FindByID(id)
	if err == nil {
		bytes, _ := json.Marshal(updatedCharacter)
		redisclient.Set("character:"+id, string(bytes), time.Minute*10)
	}

	return nil
}

// CharacterService Methods - 5 delete
func (s *CharacterService) DeleteCharacter(id string) error {
	if err := s.characterRepository.Delete(id); err != nil {
		return err
	}

	// Delete cache after removing from DB
	redisclient.Delete("character:" + id)

	return nil
}
