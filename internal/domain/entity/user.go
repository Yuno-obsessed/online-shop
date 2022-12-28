package entity

import (
	"github.com/google/uuid"
	"regexp"
	"time"
)

type User struct {
	UUID      uuid.UUID `json:"user_uuid"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Nickname  string    `json:"nickname"`
	Age       string    `json:"age"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	Salt      string    `json:"salt"`
	Purchases uint      `json:"purchases"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) Validate(action string) map[string]string {
	errors := make(map[string]string)
	switch action {
	case "login":
		reg := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		match := reg.Match([]byte(u.Email))
		if !match {
			errors["Email"] = "Please provide a valid email address"
		}
		if u.Password == "" {
			errors["Password"] = "Please enter a password"
		}
	case "register":

	}
	return errors
}
