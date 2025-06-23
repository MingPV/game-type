package dto

import "github.com/MingPV/clean-go-template/internal/entities"

func ToMonsterLootResponse(monsterLoot *entities.MonsterLoot) *MonsterLootResponse {
	return &MonsterLootResponse{
		MonsterID:   monsterLoot.MonsterID,
		ItemID:      monsterLoot.ItemID,
		DropRate:    monsterLoot.DropRate,
		QuantityMin: monsterLoot.QuantityMin,
		QuantityMax: monsterLoot.QuantityMax,
		Item:        monsterLoot.Item, // this.ItemID -> Item.ID
	}

}

func ToMonsterLootResponseList(monsterLoots []*entities.MonsterLoot) []*MonsterLootResponse {
	result := make([]*MonsterLootResponse, 0, len(monsterLoots))
	for _, o := range monsterLoots {
		result = append(result, ToMonsterLootResponse(o))
	}
	return result
}
