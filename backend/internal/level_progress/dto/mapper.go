package dto

import "github.com/MingPV/clean-go-template/internal/entities"

func ToLevelProgressResponse(level_progress *entities.LevelProgress) *LevelProgressResponse {
	return &LevelProgressResponse{
		Level:       level_progress.Level,
		ExpRequired: level_progress.ExpRequired,
	}
}

func ToLevelProgressResponseList(level_progresses []*entities.LevelProgress) []*LevelProgressResponse {
	result := make([]*LevelProgressResponse, 0, len(level_progresses))
	for _, c := range level_progresses {
		result = append(result, ToLevelProgressResponse(c))
	}
	return result
}
