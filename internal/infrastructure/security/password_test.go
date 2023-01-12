package security

import (
	"testing"
)

func TestGenerateSalt(t *testing.T) {
	salt := GenerateSalt("nickname")
	if len(salt) != 64 {
		t.Errorf("Salt wasn't created right")
	}
}

func TestCustomHash(t *testing.T) {
	salt := GenerateSalt("nickname")
	password := "some_cool_super_secure_password"
	hash := CustomHash(password, salt, 10)
	if len(hash) != 64 {
		t.Errorf("Hash wasn't generated right")
	}
}

func TestVerifyHashedPassword(t *testing.T) {
	password := "some_cool_super_secure_password"
	salt := GenerateSalt(password)
	hashedPassword := CustomHash(password, salt, 10)
	err := VerifyHashedPassword(string(hashedPassword), password, salt, 10)
	if err != nil {
		t.Errorf("%v", err)
	}
}
