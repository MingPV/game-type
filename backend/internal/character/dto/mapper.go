package dto

import "github.com/MingPV/clean-go-template/internal/entities"

func ToCharacterResponse(character *entities.Character) *CharacterResponse {
	return &CharacterResponse{
		ID:         character.ID,
		UserID:     character.UserID,
		Name:       character.Name,
		Level:      character.Level,
		CurrentExp: character.CurrentExp,
		ClassID:    character.ClassID,
		CreatedAt:  character.CreatedAt,

		Class:          character.Class,
		Status:         character.Status,
		EquipmentSlots: character.EquipmentSlots,
		Inventory:      character.Inventory,
	}
}

func ToCharacterResponseList(characters []*entities.Character) []*CharacterResponse {
	result := make([]*CharacterResponse, 0, len(characters))
	for _, c := range characters {
		result = append(result, ToCharacterResponse(c))
	}
	return result
}

// type Character struct {
// 	ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"character_id"`
// 	UserID     uuid.UUID `gorm:"type:uuid" json:"user_id"`
// 	Name       string    `json:"name"`
// 	Level      int       `json:"level"`
// 	CurrentExp int       `json:"current_exp"`
// 	ClassID    uuid.UUID `gorm:"type:uuid" json:"class_id"`
// 	StatusID   uuid.UUID `gorm:"type:uuid" json:"status_id"`
// 	CreatedAt  time.Time `json:"created_at"`

// 	User   User            `gorm:"foreignKey:UserID;references:ID"`
// 	Class  Class           `gorm:"foreignKey:ClassID;references:ID" json:"class"`
// 	Status Status          `gorm:"foreignKey:StatusID;references:CharacterID" json:"status"`
// 	Slots  []EquipmentSlot `gorm:"foreignKey:CharacterID" json:"equipment_slots"`
// 	Items  []ItemInstance  `gorm:"foreignKey:OwnerCharacterID" json:"items"`
// }
