package usecase

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/MingPV/clean-go-template/internal/class/repository"
	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/MingPV/clean-go-template/pkg/redisclient"
)

// ClassService
type ClassService struct {
	repo repository.ClassRepository
}

// Init ClassService function
func NewClassService(repo repository.ClassRepository) ClassUseCase {
	return &ClassService{repo: repo}
}

// ClassService Methods - 1 create
func (s *ClassService) CreateClass(class *entities.Class) error {
	if err := s.repo.Save(class); err != nil {
		return err
	}

	// Save to Redis cache
	bytes, _ := json.Marshal(class)
	redisclient.Set("class:"+class.ID.String(), string(bytes), time.Minute*10)

	return nil
}

// ClassService Methods - 2 find all
func (s *ClassService) FindAllClasses() ([]*entities.Class, error) {
	classes, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return classes, nil
}

// ClassService Methods - 3 find by id
func (s *ClassService) FindClassByID(character_id int) (*entities.Class, error) {

	// Check if the class is in the cache
	jsonData, err := redisclient.Get("class:" + strconv.Itoa(character_id))
	if err == nil {
		var class entities.Class
		json.Unmarshal([]byte(jsonData), &class)
		// fmt.Println("Cache hit, returning from cache")
		return &class, nil
	}

	class, err := s.repo.FindByID(character_id)
	if err != nil {
		return &entities.Class{}, err
	}

	// If not found in the cache, save it to the cache
	// fmt.Println("Cache miss saving to cache")
	bytes, _ := json.Marshal(class)
	redisclient.Set("class:"+strconv.Itoa(character_id), string(bytes), time.Minute*10)

	return class, nil
}

// ClassService Methods - 4 patch
// func (s *ClassService) PatchClass(id int, class *entities.Class) error {
// 	if class.Total <= 0 {
// 		return errors.New("total must be positive")
// 	}
// 	if err := s.repo.Patch(id, class); err != nil {
// 		return err
// 	}

// 	// Update cache after patching
// 	updatedClass, err := s.repo.FindByID(id)
// 	if err == nil {
// 		bytes, _ := json.Marshal(updatedClass)
// 		redisclient.Set("class:"+strconv.Itoa(id), string(bytes), time.Minute*10)
// 	}

// 	return nil
// }

// ClassService Methods - 5 delete
func (s *ClassService) DeleteClass(character_id int) error {
	if err := s.repo.Delete(character_id); err != nil {
		return err
	}

	// Delete cache after removing from DB
	redisclient.Delete("class:" + strconv.Itoa(character_id))

	return nil
}
