package persistence

import (
	"database/sql"
	"zusammen/internal/domain/repository"
)

type Repositories struct {
	Product repository.ProductRepository
	User    repository.UserRepository
	Cart    repository.CartRepository
	Liked   repository.LikedRepository
	Db      *sql.DB
}

func NewRepositories(conn *sql.DB) (*Repositories, error) {
	return &Repositories{
		Product: NewProductRepository(conn),
		User:    NewUserRepository(conn),
		Cart:    NewCartRepo(conn),
		Liked:   NewLikedRepo(conn),
		Db:      conn,
	}, nil
}

func (r *Repositories) Close() error {
	return r.Db.Close()
}
