package dto

import "github.com/MingPV/clean-go-template/internal/entities"

func ToRarityResponse(rarity *entities.Rarity) *RarityResponse {
	return &RarityResponse{
		ID:       rarity.ID,
		Name:     rarity.Name,
		DropRate: rarity.DropRate,
	}
}
func ToRarityResponseList(rarities []*entities.Rarity) []*RarityResponse {
	result := make([]*RarityResponse, 0, len(rarities))
	for _, o := range rarities {
		result = append(result, ToRarityResponse(o))
	}
	return result
}
