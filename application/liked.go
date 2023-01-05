package application

import (
	"github.com/google/uuid"
	"zusammen/internal/domain/entity"
	"zusammen/internal/domain/repository"
)

type likedApp struct {
	lr repository.LikedRepository
}

type LikedAppInterface interface {
	AddToLiked(liked *entity.Liked) (*entity.Liked, map[string]string)
	GetLiked(likedUuid uuid.UUID) (*entity.Liked, error)
	GetSomeLiked(userUuid uuid.UUID, limit, offset int64) ([]entity.Liked, error)
	DeleteLiked(likedUuid uuid.UUID) (uuid.UUID, error)
}

var _ LikedAppInterface = &likedApp{}

func (la *likedApp) AddToLiked(liked *entity.Liked) (*entity.Liked, map[string]string) {
	return la.lr.AddToLiked(liked)
}

func (la *likedApp) GetLiked(uuid uuid.UUID) (*entity.Liked, error) {
	return la.lr.GetLiked(uuid)
}

func (la *likedApp) GetSomeLiked(userUuid uuid.UUID, limit, offset int64) ([]entity.Liked, error) {
	return la.lr.GetSomeLiked(userUuid, limit, offset)
}

func (la *likedApp) DeleteLiked(likedUuid uuid.UUID) (uuid.UUID, error) {
	return la.lr.DeleteLiked(likedUuid)
}
