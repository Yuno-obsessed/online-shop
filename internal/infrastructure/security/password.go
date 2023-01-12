package security

import (
	"bytes"
	"crypto/sha512"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"math"
	"os"
)

// Hash is a method for passwordHashing with bcrypt, which is a
//pretty good algorithm with changeable number of interactions

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// CustomHash is a method which consists of password(plain string),
// salt(unique string for every user, is randomly generated),
// pepper(common identifier for all users, is randomly generated).

func GenerateSalt(nickname string) string {
	salt := sha512.Sum512([]byte(nickname))
	return string(salt[:])
}

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

func VerifyHashedPassword(hashedPassword string, password string, salt string, iterations int) error {
	hashNew := CustomHash(password, salt, iterations)
	if !bytes.Equal(hashNew, []byte(hashedPassword)) {
		return errors.New("Passwords do not match")
	}
	return nil
}
