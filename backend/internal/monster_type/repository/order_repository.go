package repository

import "github.com/MingPV/clean-go-template/internal/entities"

type MonsterTypeRepository interface {
	Save(monsterType *entities.MonsterType) error
	FindAll() ([]*entities.MonsterType, error)
	FindByID(id string) (*entities.MonsterType, error)
	Patch(id string, monsterType *entities.MonsterType) error
	Delete(id string) error
}
