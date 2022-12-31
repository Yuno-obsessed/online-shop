package repository

import (
	"github.com/google/uuid"
	"zusammen/internal/domain/entity"
)

type ProductRepository interface {
	PostProduct(product *entity.Product) (*entity.Product, map[string]string)
	GetProduct(uuid uuid.UUID) (*entity.Product, error)
	GetProducts(limit, offset int64) ([]entity.Product, error)
	EditProduct(product *entity.Product, uuid uuid.UUID) (*entity.Product, map[string]string)
	DeleteProduct(uuid uuid.UUID) (uuid.UUID, error)
}
