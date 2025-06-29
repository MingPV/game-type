package usecase

import "github.com/MingPV/clean-go-template/internal/entities"

type MonsterLootUseCase interface {
	FindAllMonsterLoots() ([]*entities.MonsterLoot, error)
	CreateMonsterLoot(monsterLoot *entities.MonsterLoot) (*entities.MonsterLoot, error)
	PatchMonsterLoot(monster_id string, item_id string, monsterLoot *entities.MonsterLoot) error
	DeleteMonsterLoot(monster_id string, item_id string) error
	// FindMonsterLootByID(id string) (*entities.MonsterLoot, error)
	FindMonsterLootByMonsterID(monster_id string) ([]*entities.MonsterLoot, error)
	FindMonsterLootByItemID(item_id string) ([]*entities.MonsterLoot, error)
	FindMonsterLootByMonsterIDAndItemID(monster_id string, item_id string) (*entities.MonsterLoot, error)
}
