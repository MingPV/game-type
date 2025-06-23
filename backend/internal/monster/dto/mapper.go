package dto

import "github.com/MingPV/clean-go-template/internal/entities"

func ToMonsterResponse(monster *entities.Monster) *MonsterResponse {
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
		MonsterLoots:  monster.MonsterLoots,
	}

}

func ToMonsterResponseList(monsters []*entities.Monster) []*MonsterResponse {
	result := make([]*MonsterResponse, 0, len(monsters))
	for _, o := range monsters {
		result = append(result, ToMonsterResponse(o))
	}
	return result
}

// type Monster struct {
// 	ID            uuid.UUID `gorm:"type:uuid;primaryKey" json:"monster_id"`
// 	Name          string    `json:"name"`
// 	Description   string    `json:"description"`
// 	MonsterTypeID    uuid.UUID `gorm:"type:uuid" json:"monster_type_id"`
// 	RarityID      uuid.UUID `gorm:"type:uuid" json:"rarity_id"`
// 	RequiredLevel int       `json:"required_level"`
// 	MaxStack      int       `json:"max_stack"`

// 	MonsterType  MonsterType        `gorm:"foreignKey:MonsterTypeID;references:ID" json:"monster_type"`   // this.MonsterTypeID -> MonsterType.ID
// 	Rarity    Rarity          `gorm:"foreignKey:RarityID;references:ID" json:"rarity"`        // this.RarityID -> Rarity.ID
// 	MonsterStats []MonsterLevelStat `gorm:"foreignKey:MonsterID;references:ID" json:"monster_level_stat"` // MonsterLevelStat.MonsterID -> this.ID
// }
