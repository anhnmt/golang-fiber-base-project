package repository

import (
	"sync"

	"github.com/xdorro/golang-fiber-base-project/pkg/ent"
)

var (
	once           *sync.Once
	userRepository *UserRepositoryImpl
)

type UserRepository interface {
	FindAllUsers() ([]*ent.User, error)
	FindUserByEmailAndStatusNotIn(email string, status []int) (*ent.User, error)
	FindUserByEmailAndStatusNotInAndIgnorePassword(email string, status []int) (*ent.User, error)
	FindUserByIDAndStatusNotIn(ID int64, status []int) (*ent.User, error)
	FindUserByIDAndStatusNotInAndIgnorePassword(ID int64, status []int) (*ent.User, error)
	CreateUser(*ent.User) (*ent.User, error)
	UpdateUser(*ent.User) (*ent.User, error)
	UpdateStatus(*ent.User) (*ent.User, error)
}
