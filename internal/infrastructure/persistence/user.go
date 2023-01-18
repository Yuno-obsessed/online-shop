package persistence

import (
	"database/sql"
	"errors"
	"fmt"
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
	// Check whether the email user filled in is already in use
	queryEmail := `SELECT user_uuid FROM users WHERE email=?;`
	row := r.Conn.QueryRow(queryEmail, user.Email)
	var duplicateEmail string
	_ = row.Scan(&duplicateEmail)
	if duplicateEmail != "" {
		dbErrors["email_exists"] = "this email is already taken"
	}
	// Check whether the phone number user filled in is already in use
	queryPhone := `SELECT user_uuid FROM users WHERE phone=?;`
	row = r.Conn.QueryRow(queryPhone, user.Phone)
	var duplicatePhone string
	_ = row.Scan(&duplicatePhone)
	if duplicatePhone != "" {
		dbErrors["phone_exists"] = "this phone number is already taken"
	}
	queryNickname := `SELECT user_uuid FROM users WHERE nickname=?;`
	row = r.Conn.QueryRow(queryNickname, user.Nickname)
	var duplicateNickname string
	_ = row.Scan(&duplicateNickname)
	if duplicateNickname != "" {
		dbErrors["nickname_exists"] = "this nickname is already taken"
	}
	queryImage := `SELECT user_uuid FROM users WHERE image=?;`
	row = r.Conn.QueryRow(queryImage, user.Image)
	var duplicateImage string
	_ = row.Scan(&duplicateImage)
	if duplicateImage != "" {
		dbErrors["image_exists"] = "this image is already taken"
	}
	if len(dbErrors) > 0 {
		return nil, dbErrors
	}
	// Hash the password before inserting it in database
	//user.Password = string(security.CustomHash(user.Password, user.Salt, 10))
	var err error
	user.Password, err = security.Hash(user.Password)
	query := `INSERT INTO users (user_uuid,first_name,last_name,nickname,age,email,phone,password,image,purchases,created_at,updated_at)
				VALUES (?,?,?,?,?,?,?,?,?,?,?,?);`
	_, err = r.Conn.Exec(query, user.UUID, user.FirstName, user.LastName, user.Nickname, user.Age,
		user.Email, user.Phone, user.Password, user.Image, user.Purchases, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		dbErrors["save_user"] = fmt.Sprintf("inserting user, %v", err)
		return nil, dbErrors
	}
	return user, nil
}

func (r *UserRepo) GetUser(userUuid uuid.UUID) (*entity.User, error) {
	var resUser entity.User
	query := `SELECT * FROM users WHERE user_uuid=?;`
	row := r.Conn.QueryRow(query, userUuid)
	err := row.Scan(&resUser.UUID, &resUser.FirstName, &resUser.LastName, &resUser.Nickname,
		&resUser.Age, &resUser.Email, &resUser.Phone, &resUser.Password,
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
			&resUsers[index].Phone, &resUsers[index].Password,
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
		&resUser.Age, &resUser.Email, &resUser.Phone, &resUser.Password,
		&resUser.Purchases, &resUser.CreatedAt, &resUser.UpdatedAt)
	if err != nil {
		dbErrors["get_user_by_email"] = "no such user found"
		return nil, dbErrors
	}
	//err = security.VerifyHashedPassword(resUser.Password, user.Password, resUser.Salt, 10)
	err = security.VerifyPassword(resUser.Password, user.Password)
	if err != nil {
		dbErrors["incorrect_password"] = err.Error()
		return nil, dbErrors
	}
	return &resUser, nil
}
