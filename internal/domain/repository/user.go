package repository

import (
	"github.com/google/uuid"
	"zusammen/internal/domain/entity"
)

// In this repo interface we return map when it's needed to let the user
// know his mistake at once, using templates to inject errors in html page

// By using errors we just get an error

type UserRepository interface {
	SaveUser(user *entity.User) (*entity.User, map[string]string)
	GetUser(uuid uuid.UUID) (*entity.User, error)
	GetUsers(limit, offset int64) ([]entity.User, error)
	GetUserByEmailAndPassword(user *entity.User) (*entity.User, map[string]string)
}
