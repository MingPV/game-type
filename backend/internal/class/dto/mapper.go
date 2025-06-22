package dto

import "github.com/MingPV/clean-go-template/internal/entities"

func ToClassResponse(class *entities.Class) *ClassResponse {
	return &ClassResponse{
		ID:          class.ID,
		Name:        class.Name,
		Description: class.Description,
	}
}

func ToClassResponseList(classes []*entities.Class) []*ClassResponse {
	result := make([]*ClassResponse, 0, len(classes))
	for _, c := range classes {
		result = append(result, ToClassResponse(c))
	}
	return result
}
