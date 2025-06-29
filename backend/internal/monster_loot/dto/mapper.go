package dto

import (
	"github.com/MingPV/clean-go-template/internal/entities"
	itemDTO "github.com/MingPV/clean-go-template/internal/item/dto"
)

func ToMonsterLootResponse(monsterLoot *entities.MonsterLoot) *MonsterLootResponse {
	return &MonsterLootResponse{
		MonsterID:   monsterLoot.MonsterID,
		ItemID:      monsterLoot.ItemID,
		QuantityMin: monsterLoot.QuantityMin,
		QuantityMax: monsterLoot.QuantityMax,
		Item:        *itemDTO.ToItemResponse(&monsterLoot.Item), // this.ItemID -> Item.ID
	}

}

func ToMonsterLootResponse2(monsterLoot entities.MonsterLoot) MonsterLootResponse {
	return MonsterLootResponse{
		MonsterID:   monsterLoot.MonsterID,
		ItemID:      monsterLoot.ItemID,
		QuantityMin: monsterLoot.QuantityMin,
		QuantityMax: monsterLoot.QuantityMax,
		Item:        *itemDTO.ToItemResponse(&monsterLoot.Item), // this.ItemID -> Item.ID
	}

}

func ToMonsterLootResponseList(monsterLoots []*entities.MonsterLoot) []*MonsterLootResponse {
	result := make([]*MonsterLootResponse, 0, len(monsterLoots))
	for _, o := range monsterLoots {
		result = append(result, ToMonsterLootResponse(o))
	}
	return result
}
