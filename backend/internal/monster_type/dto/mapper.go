package dto

import "github.com/MingPV/clean-go-template/internal/entities"

func ToMonsterTypeResponse(monsterType *entities.MonsterType) *MonsterTypeResponse {
	return &MonsterTypeResponse{
		ID:   monsterType.ID,
		Name: monsterType.Name,
	}
}

func ToMonsterTypeResponseList(monsterTypes []*entities.MonsterType) []*MonsterTypeResponse {
	result := make([]*MonsterTypeResponse, 0, len(monsterTypes))
	for _, o := range monsterTypes {
		result = append(result, ToMonsterTypeResponse(o))
	}
	return result
}
