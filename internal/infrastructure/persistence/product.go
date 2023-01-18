package persistence

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"zusammen/internal/domain/entity"
	"zusammen/internal/domain/repository"
)

type ProductRepo struct {
	Conn *sql.DB
}

func NewProductRepository(conn *sql.DB) *ProductRepo {
	return &ProductRepo{Conn: conn}
}

var _ repository.ProductRepository = &ProductRepo{}

func (r *ProductRepo) PostProduct(p *entity.Product) (*entity.Product, map[string]string) {
	dbErr := make(map[string]string)
	// Inserting a product to DB
	//query := `SELECT * FROM products WHERE image=?;`
	//row := r.Conn.QueryRow(query, &p.Image)
	//if row != nil {
	//	p.Image = "(1)" + p.Image
	//}
	query := `INSERT INTO products (product_uuid, product_name, description, category, image, seller, price, quantity, likes, created_at, updated_at)
				VALUES (?,?,?,?,?,?,?,?,?,?,?);`
	_, err := r.Conn.Exec(query, p.UUID, p.Name, p.Description, p.Category, p.Image, p.Seller, p.Price, p.Quantity, p.Likes, p.CreatedAt, p.UpdatedAt)
	if err != nil {
		dbErr["Error posting a product"] = err.Error()
		return nil, dbErr
	}
	return p, nil
}

func (r *ProductRepo) GetProduct(uuid uuid.UUID) (*entity.Product, error) {
	var resProduct entity.Product
	query := `SELECT * FROM products WHERE product_id=?;`
	row := r.Conn.QueryRow(query, uuid)
	err := row.Scan(&resProduct.UUID, &resProduct.Name, &resProduct.Description, &resProduct.Category, &resProduct.Image,
		&resProduct.Seller, &resProduct.Price, &resProduct.Quantity,
		&resProduct.Likes, &resProduct.CreatedAt, &resProduct.UpdatedAt)
	if err != nil {
		return &resProduct, fmt.Errorf("error in query GetProduct %v", err)
	}
	return &resProduct, nil
}

func (r *ProductRepo) GetProducts(limit, offset int64) ([]entity.Product, error) {
	var resProducts []entity.Product
	// OFFSET is a number after which goes querying up to LIMIT.
	// LIMIT 5 OFFSET 2 (3-8)
	query := `SELECT * FROM products LIMIT ? OFFSET ?;`
	rows, err := r.Conn.Query(query, limit, offset)
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("error in query GetProducts: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var prod entity.Product
		err = rows.Scan(&prod.UUID, &prod.Name, &prod.Description,
			&prod.Category, &prod.Image, &prod.Seller, &prod.Price,
			&prod.Quantity, &prod.Likes, &prod.CreatedAt, &prod.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error in query GetProducts %v", err)
		}
		resProducts = append(resProducts, prod)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error in query GetProducts %v", err)
	}
	return resProducts, nil
}

func (r *ProductRepo) EditProduct(p *entity.Product, productUuid uuid.UUID) (*entity.Product, map[string]string) {
	dbErr := make(map[string]string)
	query := `UPDATE products SET product_name=?, description=?, category=?, image=?, 
                    seller=?,price=?,quantity=?,likes=?,created_at=?,updated_at=? WHERE product_uuid=?;`
	_, err := r.Conn.Exec(query, &p.Name, &p.Description, &p.Category, &p.Image, &p.Seller, &p.Price, &p.Quantity, &p.Likes, &p.CreatedAt, &p.UpdatedAt, &productUuid)
	if err != nil {
		dbErr["error editing product"] = err.Error()
		return nil, dbErr
	}
	return p, nil
}

func (r *ProductRepo) DeleteProduct(productUuid uuid.UUID) (uuid.UUID, error) {
	query := `DELETE FROM products WHERE product_uuid=?`
	_, err := r.Conn.Exec(query, productUuid)
	if err != nil {
		return uuid.Nil, fmt.Errorf("error in query DeleteProduct %v", err)
	}
	return productUuid, nil
}
