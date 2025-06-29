package usecase

import "github.com/MingPV/clean-go-template/internal/entities"

type MonsterUseCase interface {
	FindAllMonsters() ([]*entities.Monster, error)
	CreateMonster(monster *entities.Monster) (*entities.Monster, error)
	PatchMonster(id string, monster *entities.Monster) error
	DeleteMonster(id string) error
	FindMonsterByID(id string) (*entities.Monster, error)
}
