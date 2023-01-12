package entity

import (
	"fmt"
	"github.com/google/uuid"
	"regexp"
	"strings"
)

type Product struct {
	UUID        uuid.UUID `json:"product_uuid"`
	Name        string    `json:"product_name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Image       string    `json:"image"`
	//Seller      uuid.UUID `json:"seller"`
	Seller    string `json:"seller"`
	Price     int    `json:"price"`
	Quantity  int    `json:"quantity"`
	Likes     int    `json:"likes"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (p *Product) Validate() map[string]string {
	errors := make(map[string]string)
	reg := regexp.MustCompile(`[a-zA-Z]*.`)
	match := reg.Match([]byte(p.Name))
	if !match {
		errors["Name"] = "Please enter a valid product name"
	}
	if strings.TrimSpace(p.Description) == "" {
		errors["Description"] = "Please enter a description"
	}
	//reg = regexp.MustCompile("^\\d+(\\.\\d{0,2})$")
	reg = regexp.MustCompile(`\d+`)
	match = reg.Match([]byte(fmt.Sprintf("%v", p.Price)))
	if !match {
		errors["Price"] = "Please enter a valid price"
	}
	reg = regexp.MustCompile(`\d+`)
	match = reg.Match([]byte(fmt.Sprintf("%v", p.Quantity)))
	if !match {
		errors["Quantity"] = "Please enter a valid quantity"
	}
	return errors
}
