package dto

import (
	"github.com/MingPV/clean-go-template/internal/constants"
	"github.com/MingPV/clean-go-template/internal/entities"
)

func ToStatusResponse(status *entities.Status) *StatusResponse {
	return &StatusResponse{
		CharacterID:   status.CharacterID,
		StatusPoint:   status.StatusPoint,
		AttackLevel:   status.AttackLevel,
		DefenseLevel:  status.DefenseLevel,
		HPLevel:       status.HPLevel,
		MPLevel:       status.MPLevel,
		CriticalLevel: status.CriticalLevel,
		Attack:        status.AttackLevel * constants.CAL_ATTACK,
		Defense:       status.DefenseLevel * constants.CAL_DEFENSE,
		HP:            status.HPLevel * constants.CAL_HP,
		MP:            status.MPLevel * constants.CAL_MP,
		Critical:      float32(status.CriticalLevel) * constants.CAL_CRITICAL,
	}
}

func ToStatusResponseList(statuses []*entities.Status) []*StatusResponse {
	result := make([]*StatusResponse, 0, len(statuses))
	for _, c := range statuses {
		result = append(result, ToStatusResponse(c))
	}
	return result
}
