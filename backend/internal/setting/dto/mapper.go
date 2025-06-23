package dto

import "github.com/MingPV/clean-go-template/internal/entities"

func ToSettingResponse(setting *entities.Setting) *SettingResponse {
	return &SettingResponse{
		ID:          setting.ID,
		UserID:      setting.UserID,
		MusicVolume: setting.MusicVolume,
		Language:    setting.Language,
		User:        setting.User,
	}
}

func ToSettingResponseList(settings []*entities.Setting) []*SettingResponse {
	result := make([]*SettingResponse, 0, len(settings))
	for _, o := range settings {
		result = append(result, ToSettingResponse(o))
	}
	return result
}
