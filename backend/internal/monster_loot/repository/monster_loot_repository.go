package repository

import "github.com/MingPV/clean-go-template/internal/entities"

type MonsterLootRepository interface {
	Save(monsterLoot *entities.MonsterLoot) error
	FindAll() ([]*entities.MonsterLoot, error)
	FindByID(id string) (*entities.MonsterLoot, error)
	Patch(id string, monsterLoot *entities.MonsterLoot) error
	Delete(id string) error
}
