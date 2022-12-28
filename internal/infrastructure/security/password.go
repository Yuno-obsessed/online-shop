package security

import (
	"bytes"
	"crypto/sha512"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"math"
	"os"
	"zusammen/internal/domain/entity"
)

// Hash is a method for passwordHashing with bcrypt, which is a
//pretty good algorithm with changable number if interactions

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// CustomHash is a method which consists of password(plain string),
// salt(unique string for every use, is randomly generated),
// pepper(common indentifier for all users, if randomly generated).

func CustomHash(password string, salt string, iterations int) []byte {
	iter := math.Pow(float64(2), float64(iterations))
	pepper := os.Getenv("PEPPER")
	primaryHash := sha512.Sum512([]byte(salt + password + pepper))
	finalHash := sha512.Sum512(primaryHash[:])
	for ; iter > 0; iter-- {
		finalHash = sha512.Sum512(primaryHash[:])
	}
	return finalHash[:]
}

func VerifyHashedPassword(user entity.User, password string, iterations int) error {
	hashNew := CustomHash(password, user.Salt, iterations)
	if !bytes.Equal(hashNew, []byte(user.Password)) {
		return errors.New("Passwords do not match")
	}
	return nil
}
