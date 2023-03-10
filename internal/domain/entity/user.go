package entity

import (
	"github.com/google/uuid"
	"regexp"
	"strconv"
)

type User struct {
	UUID      uuid.UUID `json:"user_uuid"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Nickname  string    `json:"nickname"`
	Age       string    `json:"age"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Image     string    `json:"image"`
	Password  string    `json:"password"`
	//Salt      string    `json:"salt"`
	Purchases uint   `json:"purchases"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (u *User) Validate(action string, secondPassword string) map[string]string {
	errors := make(map[string]string)
	switch action {

	case "login":
		reg := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		if match := reg.Match([]byte(u.Email)); !match {
			errors["Email"] = "Please provide a valid email address"
		}
		if u.Password == "" {
			errors["Password"] = "Please enter a password"
		}

	case "register":

		if u.FirstName == "" {
			errors["FirstName"] = "Please enter your first name"
		}

		if u.LastName == "" {
			errors["LastName"] = "Please enter your last name"
		}

		reg := regexp.MustCompile("[0-9]{2}")
		if match := reg.Match([]byte(u.Age)); !match {
			errors["Age"] = "Please enter a proper age"
		}
		if age, _ := strconv.Atoi(u.Age); age <= 12 {
			errors["Age"] = "Your age isn't enough to use our service"
		}

		reg = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		if match := reg.Match([]byte(u.Email)); !match {
			errors["Email"] = "Please provide a valid email address"
		}

		//123-456-7890
		//(123) 456-7890
		//123 456 7890
		//123.456.7890
		//+91 (123) 456-7890
		reg = regexp.MustCompile(`^(\+\d{1,2}\s)?\(?\d{3}\)?[\s.-]\d{3}[\s.-]\d{4}$`)
		if match := reg.Match([]byte(u.Phone)); !match {
			errors["Phone"] = "Please enter a proper phone number"
		}
		if u.Password == "" {
			errors["Password"] = "Please enter a password"
		}
		if secondPassword != u.Password {
			errors["PasswordRepeat"] = "Please enter two equal passwords"
		}
	}
	return errors
}
