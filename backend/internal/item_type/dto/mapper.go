package dto

import "github.com/MingPV/clean-go-template/internal/entities"

func ToItemTypeResponse(itemType *entities.ItemType) *ItemTypeResponse {
	return &ItemTypeResponse{
		ID:   itemType.ID,
		Name: itemType.Name,
	}
}

func ToItemTypeResponseList(itemTypes []*entities.ItemType) []*ItemTypeResponse {
	result := make([]*ItemTypeResponse, 0, len(itemTypes))
	for _, o := range itemTypes {
		result = append(result, ToItemTypeResponse(o))
	}
	return result
}
