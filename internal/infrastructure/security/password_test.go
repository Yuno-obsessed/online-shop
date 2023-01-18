package security

import (
	"fmt"
	"testing"
)

func TestGenerateSalt(t *testing.T) {
	salt := GenerateSalt("nickname")
	fmt.Println(string(salt))
	if len(salt) != 64 {
		t.Errorf("Salt wasn't created right")
	}
}

func TestCustomHash(t *testing.T) {
	salt := GenerateSalt("nickname")
	password := "some_cool_super_secure_password"
	hash := CustomHash(password, string(salt), 10)
	if len(hash) != 64 {
		t.Errorf("Hash wasn't generated right")
	}
}

func TestVerifyHashedPassword(t *testing.T) {
	password := "some_cool_super_secure_password"
	salt := string(GenerateSalt(password))
	hashedPassword := CustomHash(password, salt, 10)
	err := VerifyHashedPassword(string(hashedPassword), password, salt, 10)
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestHash(t *testing.T) {
	hashedPassword, err := Hash("some_cool_super_secure_password")
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Println(hashedPassword)
	fmt.Println(len(hashedPassword))
	if len(hashedPassword) != 60 {
		t.Errorf("Generated hash doesn't seem to be right")
	}
}

func TestVerifyPassword(t *testing.T) {
	hashedPassword, err := Hash("some_cool_super_secure_password")
	if err != nil {
		t.Errorf("%v", err)
	}
	if len(hashedPassword) != 60 {
		t.Errorf("Generated hash doesn't seem to be right")
	}
	repeatPassword := "some_cool_super_secure_password"
	err = VerifyPassword(hashedPassword, repeatPassword)
	if err != nil {
		t.Errorf("%v", err)
	}
}
