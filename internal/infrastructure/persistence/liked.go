package persistence

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"zusammen/internal/domain/entity"
)

type LikedRepo struct {
	Conn *sql.DB
}

func NewLikedRepo(conn *sql.DB) *LikedRepo {
	return &LikedRepo{Conn: conn}
}

func (lr *LikedRepo) AddToLiked(l *entity.Liked) (*entity.Liked, map[string]string) {
	dbErr := make(map[string]string)
	// Making sure not to like one product more then one time
	// Should think about how to make it work for every user separately
	query := `SELECT * FROM likes WHERE product_uuid=?;`
	match := lr.Conn.QueryRow(query, l.ProductUUID)
	if match != nil {
		dbErr["Error adding to liked"] = "This product is already liked"
		return nil, dbErr
	}
	query = `INSERT INTO likes (liked_uuid, user_uuid, product_uuid)
				VALUES (?,?,?);`
	_, err := lr.Conn.Exec(query, l.UUID, l.UserUUID, l.ProductUUID)
	if err != nil {
		dbErr["Error adding to liked"] = err.Error()
		return nil, dbErr
	}
	return l, nil
}
func (lr *LikedRepo) GetLiked(likedUuid uuid.UUID) (*entity.Liked, error) {
	var resLiked entity.Liked
	query := `SELECT * FROM likes WHERE liked_uuid=?;`
	row := lr.Conn.QueryRow(query, likedUuid)
	err := row.Scan(&resLiked.UUID, &resLiked.UserUUID, resLiked.ProductUUID)
	if err != nil {
		return nil, fmt.Errorf("GetLiked query scan error %v", err)
	}
	return &resLiked, nil
}
func (lr *LikedRepo) GetSomeLiked(userUuid uuid.UUID, limit, offset int64) ([]entity.Liked, error) {
	var resSomeLiked []entity.Liked
	query := `SELECT * FROM likes WHERE user_uuid=? LIMIT ? OFFSET ?;`
	res, err := lr.Conn.Query(query, userUuid, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("GetSomeLiked query scan error %v", err)
	}
	res.Close()
	// Find a solution to a problem where if there are less than 5 products, next page
	// gives an error instead of uploading 3 products
	for res.Next() {
		var liked entity.Liked
		err = res.Scan(&liked.UUID, &liked.UserUUID, &liked.ProductUUID)
		if err != nil {
			return nil, fmt.Errorf("GetSomeLiked query scan error %v", err)
		}
		resSomeLiked = append(resSomeLiked, liked)
	}
	if err = res.Err(); err != nil {
		return nil, fmt.Errorf("GetSomeLiked query scan error %v", err)
	}
	return resSomeLiked, nil
}
func (lr *LikedRepo) DeleteLiked(likedUuid uuid.UUID) (uuid.UUID, error) {
	query := `DELETE FROM liked WHERE liked_uuid=?;`
	_, err := lr.Conn.Exec(query, likedUuid)
	if err != nil {
		return uuid.Nil, fmt.Errorf("DeleteLiked query error %v", err)
	}
	return likedUuid, nil
}
