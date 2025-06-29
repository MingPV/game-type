package dto

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	monsterLootDTO "github.com/MingPV/clean-go-template/internal/monster_loot/dto"
)

func ToMonsterResponse(monster *entities.Monster) *MonsterResponse {

	monster_loots := make([]monsterLootDTO.MonsterLootResponse, 0, len(monster.MonsterLoots))
	for _, m := range monster.MonsterLoots {
		monster_loots = append(monster_loots, monsterLootDTO.ToMonsterLootResponse2(m))
	}

	return &MonsterResponse{
		ID:            monster.ID,
		Name:          monster.Name,
		Description:   monster.Description,
		Level:         monster.Level,
		HP:            monster.HP,
		Attack:        monster.Attack,
		Defense:       monster.Defense,
		ExpReward:     monster.ExpReward,
		GoldReward:    monster.GoldReward,
		MonsterTypeID: monster.MonsterTypeID,
		MonsterType:   monster.MonsterType,
		MonsterLoots:  monster_loots,
	}

}

func ToMonsterResponseList(monsters []*entities.Monster) []*MonsterResponse {
	result := make([]*MonsterResponse, 0, len(monsters))
	for _, o := range monsters {
		result = append(result, ToMonsterResponse(o))
	}
	return result
}
