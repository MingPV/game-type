package dto

import "github.com/MingPV/clean-go-template/internal/entities"

func ToStatusResponse(status *entities.Status) *StatusResponse {
	return &StatusResponse{
		CharacterID: status.CharacterID,
		StatusPoint: status.StatusPoint,
		STR:         status.STR,
		AGI:         status.AGI,
		INT:         status.INT,
		DEX:         status.DEX,
		VIT:         status.VIT,
		LUK:         status.LUK,
	}
}

func ToStatusResponseList(statuses []*entities.Status) []*StatusResponse {
	result := make([]*StatusResponse, 0, len(statuses))
	for _, c := range statuses {
		result = append(result, ToStatusResponse(c))
	}
	return result
}
