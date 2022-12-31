package application

import (
	"github.com/google/uuid"
	"zusammen/internal/domain/entity"
	"zusammen/internal/domain/repository"
)

type userApp struct {
	us repository.UserRepository
}

var _ UserAppInterface = &userApp{}

type UserAppInterface interface {
	SaveUser(user *entity.User) (*entity.User, map[string]string)
	GetUser(uuid uuid.UUID) (*entity.User, error)
	GetUsers(limit, offset int64) ([]entity.User, error)
	// EditUser(user *entity.User, uuid uuid.UUID) (*entity.Product, map[string]string)
	GetUserByEmailAndPassword(user *entity.User) (*entity.User, map[string]string)
}

func (u *userApp) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	return u.us.SaveUser(user)
}

func (u *userApp) GetUser(uuid uuid.UUID) (*entity.User, error) {
	return u.us.GetUser(uuid)
}

func (u *userApp) GetUsers(limit, offset int64) ([]entity.User, error) {
	return u.us.GetUsers(limit, offset)
}

func (u *userApp) GetUserByEmailAndPassword(user *entity.User) (*entity.User, map[string]string) {
	return u.us.GetUserByEmailAndPassword(user)
}
