package repository

import (
	"github.com/google/uuid"
	"zusammen/internal/domain/entity"
)

type LikedRepository interface {
	AddToLiked(liked *entity.Liked) (*entity.Liked, map[string]string)
	GetLiked(likedUuid uuid.UUID) (*entity.Liked, error)
	GetSomeLiked(userUuid uuid.UUID, limit, offset int64) ([]entity.Liked, error)
	DeleteLiked(likedUuid uuid.UUID) (uuid.UUID, error)
}
