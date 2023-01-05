package auth

import (
	"strings"
	"testing"
)

func TestJWT_GenerateToken(t *testing.T) {
	jwt := JWT{Secret: "secret_key"}
	token, err := jwt.GenerateToken("sanity_nil")
	if err != nil {
		t.Errorf("Error generating JWT, %v", err)
	}
	if !strings.Contains(token, ".") {
		t.Errorf("Invalid JWT format, %v", err)
	}
}

func TestJWT_ValidateToken(t *testing.T) {
	jwt := JWT{Secret: "secret_key"}
	username := "sanity_nil"
	token, err := jwt.GenerateToken(username)
	if err != nil {
		t.Errorf("Error generating JWT, %v", err)
	}
	if !strings.Contains(token, ".") {
		t.Errorf("Invalid JWT format, %v", err)
	}
	claims, err := jwt.ValidateToken(token)
	if err != nil {
		t.Errorf("Error validating JWT, %v", err)
	}
	if username != claims.Username {
		t.Errorf("Invalid claims, %v", err)
	}
}
