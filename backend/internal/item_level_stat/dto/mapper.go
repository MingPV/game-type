package dto

import "github.com/MingPV/clean-go-template/internal/entities"

func ToItemLevelStatResponse(itemLevelStat *entities.ItemLevelStat) *ItemLevelStatResponse {
	return &ItemLevelStatResponse{
		ItemID:   itemLevelStat.ItemID,
		Level:    itemLevelStat.Level,
		BonusSTR: itemLevelStat.BonusSTR,
		BonusAGI: itemLevelStat.BonusAGI,
		BonusINT: itemLevelStat.BonusINT,
		BonusDEX: itemLevelStat.BonusDEX,
		BonusVIT: itemLevelStat.BonusVIT,
		BonusLUK: itemLevelStat.BonusLUK,
	}
}

func ToItemLevelStatResponseList(itemLevelStats []*entities.ItemLevelStat) []*ItemLevelStatResponse {
	result := make([]*ItemLevelStatResponse, 0, len(itemLevelStats))
	for _, o := range itemLevelStats {
		result = append(result, ToItemLevelStatResponse(o))
	}
	return result
}
