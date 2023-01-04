package repository

import (
	"github.com/google/uuid"
	"zusammen/internal/domain/entity"
)

type CartRepository interface {
	AddToCart(c *entity.Carts) (*entity.Carts, map[string]string)
	GetCart(cartUuid uuid.UUID) (*entity.Carts, error)
	GetCarts(userUuid uuid.UUID, limit, offset int64) ([]entity.Carts, error)
	//EditCart(cart *entity.Carts, uuid uuid.UUID) (*entity.Carts, map[string]string)
	DeleteCart(cartUuid uuid.UUID) (uuid.UUID, error)
}
