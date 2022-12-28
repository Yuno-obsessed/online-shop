package persistence

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"zusammen/internal/domain/entity"
	"zusammen/internal/domain/repository"
	"zusammen/internal/infrastructure/security"
)

type UserRepo struct {
	Conn *sql.DB
}

func NewUserRepository(conn *sql.DB) *UserRepo {
	return &UserRepo{Conn: conn}
}

var _ repository.UserRepository = &UserRepo{}

func (r *UserRepo) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	dbErrors := make(map[string]string)
	userUuid := uuid.New()
	query_email := `SELECT user_uuid FROM users WHERE email=?;`
	duplicate_email := r.Conn.QueryRow(query_email, user.Email)
	err := duplicate_email.Err()
	if err != nil {
		dbErrors["email_exists"] = "this email is already taken"
	}
	query_phone := `SELECT user_uuid FROM users WHERE phone=?;`
	duplicate_phone := r.Conn.QueryRow(query_phone, user.Phone)
	err = duplicate_phone.Err()
	if err != nil {
		dbErrors["phone_exists"] = "this phone number is already taken"
	}
	if dbErrors != nil {
		return nil, dbErrors
	}
	query := `INSERT INTO users (user_uuid,first_name,last_name,nickname,age,email,phone,password,salt,created_at,updated_at)
				VALUES (?,?,?,?,?,?,?,?,?,?);`
	_, err = r.Conn.Exec(query, &userUuid, &user.FirstName, &user.LastName, &user.Nickname, &user.Age,
		&user.Email, &user.Phone, &user.Password, &user.Salt, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		dbErrors["save_user"] = "inserting operation"
		return nil, dbErrors
	}
	return user, nil
}

func (r *UserRepo) GetUser(userUuid uuid.UUID) (*entity.User, error) {
	var resUser entity.User
	query := `SELECT * FROM users WHERE user_uuid=?;`
	row := r.Conn.QueryRow(query, userUuid)
	err := row.Scan(&resUser.UUID, &resUser.FirstName, &resUser.LastName, &resUser.Nickname,
		&resUser.Age, &resUser.Email, &resUser.Phone, &resUser.Password, &resUser.Salt,
		&resUser.Purchases, &resUser.CreatedAt, &resUser.UpdatedAt)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &resUser, nil
}

func (r *UserRepo) GetUsers(limit, offset int64) ([]entity.User, error) {
	var resUsers []entity.User
	// Say you want to get 5 artists, but not the first five. You want to get rows 3 through 8.
	// You’ll want to add an OFFSET of 2 to skip the first two rows
	query := `SELECT * FROM users LIMIT ? OFFSET ?;`
	rows, err := r.Conn.Query(query, limit, offset)
	if err != nil {
		return nil, errors.New("users in limit not found")
	}
	defer rows.Close()
	index := 0
	for rows.Next() {
		err = rows.Scan(&resUsers[index].UUID, &resUsers[index].FirstName, &resUsers[index].LastName,
			&resUsers[index].Nickname, &resUsers[index].Age, &resUsers[index].Email,
			&resUsers[index].Phone, &resUsers[index].Password, &resUsers[index].Salt,
			&resUsers[index].Purchases, &resUsers[index].CreatedAt, &resUsers[index].UpdatedAt)
		if err != nil {
			return resUsers, err
		}
		index++
	}
	return resUsers, nil
}

func (r *UserRepo) GetUserByEmailAndPassword(user *entity.User) (*entity.User, map[string]string) {
	var resUser entity.User
	dbErrors := make(map[string]string)
	query := `SELECT * FROM users WHERE email=?;`
	row := r.Conn.QueryRow(query, user.Email)
	err := row.Scan(&resUser.UUID, &resUser.FirstName, &resUser.LastName, &resUser.Nickname,
		&resUser.Age, &resUser.Email, &resUser.Phone, &resUser.Password, &resUser.Salt,
		&resUser.Purchases, &resUser.CreatedAt, &resUser.UpdatedAt)
	if err != nil {
		dbErrors["get_user_by_email"] = "no such user found"
		return nil, dbErrors
	}
	err = security.VerifyHashedPassword(resUser, user.Password, 10)
	if err != nil {
		dbErrors["incorrect_password"] = err.Error()
		return nil, dbErrors
	}
	return &resUser, nil
}
