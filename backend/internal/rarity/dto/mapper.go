package dto

import "github.com/MingPV/clean-go-template/internal/entities"

func ToRarityResponse(rarity *entities.Rarity) *RarityResponse {
	return &RarityResponse{
		ID:       rarity.ID,
		Name:     rarity.Name,
		DropRate: rarity.DropRate,
		Items:    rarity.Items,
	}
}

// type Rarity struct {
// 	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"rarity_id"`
// 	Name     string    `json:"name"`
// 	DropRate float64   `json:"drop_rate"`

// 	Items []Item `gorm:"foreignKey:RarityID;references:ID" json:"items"` // Item.RarityID -> this.ID
// }

func ToRarityResponseList(rarities []*entities.Rarity) []*RarityResponse {
	result := make([]*RarityResponse, 0, len(rarities))
	for _, o := range rarities {
		result = append(result, ToRarityResponse(o))
	}
	return result
}
