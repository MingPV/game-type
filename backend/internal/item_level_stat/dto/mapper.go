package dto

import (
	"encoding/json"
	"fmt"

	"github.com/MingPV/clean-go-template/internal/entities"
)

func ToItemLevelStatResponse(itemLevelStat *entities.ItemLevelStat) *ItemLevelStatResponse {

	// bonus_stat := itemLevelStat.BonusStat

	fmt.Println("ItemLevelStatResponse: ", itemLevelStat.BonusStat)

	var bonusStatMap map[string]entities.BonusStat
	err := json.Unmarshal([]byte(itemLevelStat.BonusStat), &bonusStatMap)
	if err != nil {
		bonusStatMap = make(map[string]entities.BonusStat)
	}

	fmt.Println("ItemLevelStatResponse: ", bonusStatMap)

	return &ItemLevelStatResponse{
		ItemID:    itemLevelStat.ItemID,
		BonusStat: bonusStatMap,
	}
}

func ToItemLevelStatResponseList(itemLevelStats []*entities.ItemLevelStat) []*ItemLevelStatResponse {
	result := make([]*ItemLevelStatResponse, 0, len(itemLevelStats))
	for _, o := range itemLevelStats {
		result = append(result, ToItemLevelStatResponse(o))
	}
	return result
}
