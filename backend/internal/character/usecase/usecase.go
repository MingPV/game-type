package usecase

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/MingPV/clean-go-template/internal/character/repository"
	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/MingPV/clean-go-template/pkg/redisclient"
)

// CharacterService
type CharacterService struct {
	repo repository.CharacterRepository
}

// Init CharacterService function
func NewCharacterService(repo repository.CharacterRepository) CharacterUseCase {
	return &CharacterService{repo: repo}
}

// CharacterService Methods - 1 create
func (s *CharacterService) CreateCharacter(character *entities.Character) error {
	if err := s.repo.Save(character); err != nil {
		return err
	}

	// Save to Redis cache
	bytes, _ := json.Marshal(character)
	redisclient.Set("character:"+character.ID.String(), string(bytes), time.Minute*10)

	return nil
}

// CharacterService Methods - 2 find all
func (s *CharacterService) FindAllCharacters() ([]*entities.Character, error) {
	characters, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return characters, nil
}

// CharacterService Methods - 3 find by id
func (s *CharacterService) FindCharacterByID(id int) (*entities.Character, error) {

	// Check if the character is in the cache
	jsonData, err := redisclient.Get("character:" + strconv.Itoa(id))
	if err == nil {
		var character entities.Character
		json.Unmarshal([]byte(jsonData), &character)
		// fmt.Println("Cache hit, returning from cache")
		return &character, nil
	}

	character, err := s.repo.FindByID(id)
	if err != nil {
		return &entities.Character{}, err
	}

	// If not found in the cache, save it to the cache
	// fmt.Println("Cache miss saving to cache")
	bytes, _ := json.Marshal(character)
	redisclient.Set("character:"+strconv.Itoa(id), string(bytes), time.Minute*10)

	return character, nil
}

// CharacterService Methods - 4 patch
// func (s *CharacterService) PatchCharacter(id int, character *entities.Character) error {
// 	if character.Total <= 0 {
// 		return errors.New("total must be positive")
// 	}
// 	if err := s.repo.Patch(id, character); err != nil {
// 		return err
// 	}

// 	// Update cache after patching
// 	updatedCharacter, err := s.repo.FindByID(id)
// 	if err == nil {
// 		bytes, _ := json.Marshal(updatedCharacter)
// 		redisclient.Set("character:"+strconv.Itoa(id), string(bytes), time.Minute*10)
// 	}

// 	return nil
// }

// CharacterService Methods - 5 delete
func (s *CharacterService) DeleteCharacter(id int) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}

	// Delete cache after removing from DB
	redisclient.Delete("character:" + strconv.Itoa(id))

	return nil
}
