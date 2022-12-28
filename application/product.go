package application

import (
	"github.com/google/uuid"
	"zusammen/internal/domain/entity"
	"zusammen/internal/domain/repository"
)

type productApp struct {
	pr repository.ProductRepository
}

var _ ProductAppInterface = &productApp{}

type ProductAppInterface interface {
	PostProduct(product *entity.Product) (uuid.UUID, error)
	GetProduct(uuid uuid.UUID) (*entity.Product, error)
	GetProducts(limit, offset int64) ([]entity.Product, error)
	EditProduct(product *entity.Product, uuid uuid.UUID) (uuid.UUID, error)
	DeleteProduct(uuid uuid.UUID) (uuid.UUID, error)
}

func (p *productApp) PostProduct(product *entity.Product) (uuid.UUID, error) {
	return p.pr.PostProduct(product)
}

func (p *productApp) GetProduct(uuid uuid.UUID) (*entity.Product, error) {
	return p.pr.GetProduct(uuid)
}

func (p *productApp) GetProducts(limit, offset int64) ([]entity.Product, error) {
	return p.pr.GetProducts(limit, offset)
}

func (p *productApp) EditProduct(product *entity.Product, uuid uuid.UUID) (uuid.UUID, error) {
	return p.pr.EditProduct(product, uuid)
}

func (p *productApp) DeleteProduct(uuid uuid.UUID) (uuid.UUID, error) {
	return p.pr.DeleteProduct(uuid)
}
