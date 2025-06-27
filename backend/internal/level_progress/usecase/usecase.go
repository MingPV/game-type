package usecase

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/MingPV/clean-go-template/internal/entities"
	"github.com/MingPV/clean-go-template/internal/level_progress/repository"
	"github.com/MingPV/clean-go-template/pkg/redisclient"
)

// LevelProgressService
type LevelProgressService struct {
	repo repository.LevelProgressRepository
}

// Init LevelProgressService function
func NewLevelProgressService(repo repository.LevelProgressRepository) LevelProgressUseCase {
	return &LevelProgressService{repo: repo}
}

// LevelProgressService Methods - 1 create
func (s *LevelProgressService) CreateLevelProgress(level_progress *entities.LevelProgress) error {
	if err := s.repo.Save(level_progress); err != nil {
		return err
	}

	// Save to Redis cache
	bytes, _ := json.Marshal(level_progress)
	redisclient.Set("level_progress:"+strconv.Itoa(level_progress.Level), string(bytes), time.Minute*10)

	return nil
}

// LevelProgressService Methods - 2 find all
func (s *LevelProgressService) FindAllLevelProgresses() ([]*entities.LevelProgress, error) {
	level_progresses, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return level_progresses, nil
}

// LevelProgressService Methods - 3 find by level
func (s *LevelProgressService) FindLevelProgressByLevel(level int) (*entities.LevelProgress, error) {

	// Check if the level_progress is in the cache
	jsonData, err := redisclient.Get("level_progress:" + strconv.Itoa(level))
	if err == nil {
		var level_progress entities.LevelProgress
		json.Unmarshal([]byte(jsonData), &level_progress)
		// fmt.Println("Cache hit, returning from cache")
		return &level_progress, nil
	}

	level_progress, err := s.repo.FindByLevel(level)
	if err != nil {
		return &entities.LevelProgress{}, err
	}

	// If not found in the cache, save it to the cache
	// fmt.Println("Cache miss saving to cache")
	bytes, _ := json.Marshal(level_progress)
	redisclient.Set("level_progress:"+strconv.Itoa(level), string(bytes), time.Minute*10)

	return level_progress, nil
}

// LevelProgressService Methods - 4 patch
func (s *LevelProgressService) PatchLevelProgress(level int, level_progress *entities.LevelProgress) error {

	if err := s.repo.Patch(level, level_progress); err != nil {
		return err
	}

	// Update cache after patching
	updatedLevelProgress, err := s.repo.FindByLevel(level)
	if err == nil {
		bytes, _ := json.Marshal(updatedLevelProgress)
		redisclient.Set("level_progress:"+strconv.Itoa(level), string(bytes), time.Minute*10)
	}

	return nil
}

// LevelProgressService Methods - 5 delete
func (s *LevelProgressService) DeleteLevelProgress(level int) error {
	if err := s.repo.Delete(level); err != nil {
		return err
	}

	// Delete cache after removing from DB
	redisclient.Delete("level_progress:" + strconv.Itoa(level))

	return nil
}
