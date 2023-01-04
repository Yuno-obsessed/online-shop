package application

import (
	"github.com/google/uuid"
	"zusammen/internal/domain/entity"
	"zusammen/internal/domain/repository"
)

type cartApp struct {
	cr repository.CartRepository
}

type CartAppInterface interface {
	AddToCart(c *entity.Carts) (*entity.Carts, map[string]string)
	GetCart(cartUuid uuid.UUID) (*entity.Carts, error)
	GetCarts(userUuid uuid.UUID, limit, offset int64) ([]entity.Carts, error)
	//EditCart(cart *entity.Carts, uuid uuid.UUID) (*entity.Carts, map[string]string)
	DeleteCart(cartUuid uuid.UUID) (uuid.UUID, error)
}

var _ repository.CartRepository = &cartApp{}

func (ca *cartApp) AddToCart(c *entity.Carts) (*entity.Carts, map[string]string) {
	return ca.cr.AddToCart(c)
}

func (ca *cartApp) GetCart(cartUuid uuid.UUID) (*entity.Carts, error) {
	return ca.cr.GetCart(cartUuid)
}

func (ca *cartApp) GetCarts(userUuid uuid.UUID, limit, offset int64) ([]entity.Carts, error) {
	return ca.cr.GetCarts(userUuid, limit, offset)
}

func (ca *cartApp) DeleteCart(cartUuid uuid.UUID) (uuid.UUID, error) {
	return ca.cr.DeleteCart(cartUuid)
}
