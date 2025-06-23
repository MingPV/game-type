package usecase

import "github.com/MingPV/clean-go-template/internal/entities"

type MonsterTypeUseCase interface {
	FindAllMonsterTypes() ([]*entities.MonsterType, error)
	CreateMonsterType(monsterType *entities.MonsterType) error
	PatchMonsterType(id string, monsterType *entities.MonsterType) error
	DeleteMonsterType(id string) error
	FindMonsterTypeByID(id string) (*entities.MonsterType, error)
}
