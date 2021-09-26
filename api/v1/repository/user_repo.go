package repository

import (
	"context"
	"log"
	"sync"

	"github.com/xdorro/golang-fiber-base-project/pkg/ent"
	"github.com/xdorro/golang-fiber-base-project/pkg/ent/user"
)

type UserRepositoryImpl struct {
	client *ent.Client
	ctx    context.Context
}

func NewUserRepository(client *ent.Client, ctx context.Context) UserRepository {
	if userRepository == nil {
		once = &sync.Once{}

		once.Do(func() {
			userRepository = &UserRepositoryImpl{
				client: client,
				ctx:    ctx,
			}

			log.Println("Create new UserRepository")
		})
	}

	return userRepository
}

func (repo *UserRepositoryImpl) FindAllUsers() ([]*ent.User, error) {
	users := make([]*ent.User, 0)

	users, err := repo.client.User.
		Query().
		Select(user.FieldID, user.FieldName, user.FieldEmail, user.FieldStatus, user.FieldCreatedAt, user.FieldUpdatedAt).
		Where(user.StatusIn(1, 2)).
		All(repo.ctx)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (repo *UserRepositoryImpl) FindUserByEmailAndStatusNotIn(email string, status []int) (*ent.User, error) {
	findUser := new(ent.User)

	findUser, err := repo.client.User.
		Query().
		Where(user.StatusNotIn(status...), user.Email(email)).
		First(repo.ctx)

	if err != nil {
		return nil, err
	}

	return findUser, nil
}

func (repo *UserRepositoryImpl) FindUserByEmailAndStatusNotInAndIgnorePassword(email string, status []int) (*ent.User, error) {
	findUser := new(ent.User)

	findUser, err := repo.client.User.
		Query().
		Select(user.FieldID, user.FieldName, user.FieldEmail, user.FieldStatus, user.FieldCreatedAt, user.FieldUpdatedAt).
		Where(user.StatusNotIn(status...), user.Email(email)).
		First(repo.ctx)

	if err != nil {
		return nil, err
	}

	return findUser, nil
}

func (repo *UserRepositoryImpl) FindUserByIDAndStatusNotIn(ID int64, status []int) (*ent.User, error) {
	obj := new(ent.User)

	obj, err := repo.client.User.
		Query().
		Where(user.StatusNotIn(status...), user.ID(ID)).
		First(repo.ctx)

	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (repo *UserRepositoryImpl) FindUserByIDAndStatusNotInAndIgnorePassword(ID int64, status []int) (*ent.User, error) {
	obj := new(ent.User)

	obj, err := repo.client.User.
		Query().
		Select(user.FieldID, user.FieldName, user.FieldEmail, user.FieldStatus, user.FieldCreatedAt, user.FieldUpdatedAt).
		Where(user.StatusNotIn(status...), user.ID(ID)).
		First(repo.ctx)

	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (repo *UserRepositoryImpl) CreateUser(e *ent.User) (*ent.User, error) {
	obj, err := repo.client.User.
		Create().
		SetName(e.Name).
		SetEmail(e.Email).
		SetPassword(e.Password).
		SetStatus(e.Status).
		Save(repo.ctx)

	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (repo *UserRepositoryImpl) UpdateUser(e *ent.User) (*ent.User, error) {
	obj, err := repo.client.User.
		UpdateOneID(e.ID).
		SetName(e.Name).
		SetEmail(e.Email).
		SetStatus(e.Status).
		Save(repo.ctx)

	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (repo *UserRepositoryImpl) UpdateStatus(e *ent.User) (*ent.User, error) {
	obj, err := repo.client.User.
		UpdateOneID(e.ID).
		SetStatus(e.Status).
		Save(repo.ctx)

	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (repo *UserRepositoryImpl) UpdatePassword(e *ent.User) (*ent.User, error) {
	obj, err := repo.client.User.
		UpdateOneID(e.ID).
		SetPassword(e.Password).
		Save(repo.ctx)

	if err != nil {
		return nil, err
	}

	return obj, nil
}
