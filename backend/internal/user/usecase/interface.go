package usecase

import "github.com/MingPV/clean-go-template/internal/entities"

// What UserService can do
type UserUseCase interface {
	Register(user *entities.User) error
	Login(email string, password string) (string, *entities.User, error)
	LoginWithUsername(username string, password string) (string, *entities.User, error)
	FindUserByID(id string) (*entities.User, error)
	FindUserByEmail(email string) (*entities.User, error)
	FindUserByUsername(username string) (*entities.User, error)
	FindAllUsers() ([]*entities.User, error)
}
