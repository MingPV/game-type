package repository

import "github.com/MingPV/clean-go-template/internal/entities"

type MonsterLootRepository interface {
	Save(monsterLoot *entities.MonsterLoot) error
	FindAll() ([]*entities.MonsterLoot, error)
	FindByMonsterID(monster_id string) ([]*entities.MonsterLoot, error)
	FindByItemID(item_id string) ([]*entities.MonsterLoot, error)
	FindByMonsterIDAndItemID(monster_id string, item_id string) (*entities.MonsterLoot, error)
	Patch(monster_id string, item_id string, monsterLoot *entities.MonsterLoot) error
	Delete(monster_id string, item_id string) error
}
