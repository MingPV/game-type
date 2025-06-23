package usecase

import "github.com/MingPV/clean-go-template/internal/entities"

type MonsterLootUseCase interface {
	FindAllMonsterLoots() ([]*entities.MonsterLoot, error)
	CreateMonsterLoot(monsterLoot *entities.MonsterLoot) error
	PatchMonsterLoot(id string, monsterLoot *entities.MonsterLoot) error
	DeleteMonsterLoot(id string) error
	FindMonsterLootByID(id string) (*entities.MonsterLoot, error)
}
