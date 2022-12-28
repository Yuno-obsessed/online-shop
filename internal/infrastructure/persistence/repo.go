package persistence

import (
	"database/sql"
	"fmt"
	"zusammen/internal/domain/repository"
	"zusammen/internal/infrastructure/config"
)

type Repositories struct {
	Product repository.ProductRepository
	User    repository.UserRepository
	Db      *sql.DB
}

func NewRepositories(config *config.DatabaseConfig) (*Repositories, error) {
	conn, err := sql.Open(config.Driver, fmt.Sprintf(
		"%s:%s@tcp(127.0.0.1:%s)/%s",
		config.Username,
		config.Password,
		config.Port,
		config.Database))
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(); err != nil {
		return nil, err
	}
	return &Repositories{
		Product: NewProductRepository(conn),
		User:    NewUserRepository(conn),
		Db:      conn,
	}, nil
}

func (r *Repositories) Close() error {
	return r.Db.Close()
}
