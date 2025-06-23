package repository

import "github.com/MingPV/clean-go-template/internal/entities"

type MonsterRepository interface {
	Save(monster *entities.Monster) error
	FindAll() ([]*entities.Monster, error)
	FindByID(id string) (*entities.Monster, error)
	Patch(id string, monster *entities.Monster) error
	Delete(id string) error
}
