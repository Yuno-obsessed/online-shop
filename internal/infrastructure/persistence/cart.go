package persistence

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"zusammen/internal/domain/entity"
	"zusammen/internal/domain/repository"
)

type CartRepo struct {
	Conn *sql.DB
}

func NewCartRepo(conn *sql.DB) *CartRepo {
	return &CartRepo{Conn: conn}
}

var _ repository.CartRepository = &CartRepo{}

func (cr *CartRepo) AddToCart(c *entity.Carts) (*entity.Carts, map[string]string) {
	dbErr := make(map[string]string)
	c.UUID = uuid.New()
	// Checking if this product is already in a cart
	query := `SELECT * FROM carts WHERE product_uuid=?;`
	match := cr.Conn.QueryRow(query, c.ProductUUID)
	if match != nil {
		dbErr["Error adding to cart"] = "This product already is in a cart"
		return c, dbErr
	}
	query = `INSERT INTO carts (cart_uuid, user_uuid, product_uuid, quantity)
				VALUES (?,?,?,?);`
	_, err := cr.Conn.Exec(query, c.UUID, c.UserUUID, c.ProductUUID, c.Quantity)
	if err != nil {
		dbErr["Error adding to cart"] = err.Error()
		return c, dbErr
	}
	return c, dbErr
}

func (cr *CartRepo) GetCart(cartUuid uuid.UUID) (*entity.Carts, error) {
	var resCart entity.Carts
	query := `SELECT * FROM carts WHERE cart_uuid=?;`
	res := cr.Conn.QueryRow(query, cartUuid)
	err := res.Scan(&resCart.UUID, &resCart.UserUUID, &resCart.ProductUUID, &resCart.Quantity)
	if err != nil {
		return nil, fmt.Errorf("GetCart query scan error %v", err)
	}
	return &resCart, nil
}

func (cr *CartRepo) GetCarts(userUuid uuid.UUID, limit, offset int64) ([]entity.Carts, error) {
	var resCarts []entity.Carts
	query := `SELECT * FROM carts WHERE user_uuid=? LIMIT ? OFFSET ?;`
	res, err := cr.Conn.Query(query, userUuid, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("GetCarts query error %v", err)
	}
	defer res.Close()
	for res.Next() {
		var cart entity.Carts
		err = res.Scan(&cart.UUID, &cart.UserUUID, &cart.ProductUUID, &cart.Quantity)
		if err != nil {
			return nil, fmt.Errorf("GetCarts query scan error %v", err)
		}
		resCarts = append(resCarts, cart)
	}
	if err = res.Err(); err != nil {
		return nil, fmt.Errorf("GetCarts query scan error %v", err)
	}
	return resCarts, nil
}

func (cr *CartRepo) DeleteCart(cartUuid uuid.UUID) (uuid.UUID, error) {
	query := `DELETE FROM carts WHERE cart_uuid=?;`
	_, err := cr.Conn.Exec(query, cartUuid)
	if err != nil {
		return uuid.Nil, fmt.Errorf("DeleteCart exec error %v", err)
	}
	return cartUuid, nil
}
